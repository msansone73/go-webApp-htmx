package scrapping

import (
	"fmt"
	"log"
	"main/model"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type StockInfo struct {
	Stock string
	Date time.Time
	Price float64
	Dividends []model.Dividend
}

const (
	Investidor10AcoesUrl = "https://investidor10.com.br/acoes/"
	Investidor10FiiUrl = "https://investidor10.com.br/fundos-imobiliarios/"
	Investidor10Url = "https://investidor10.com.br/%s/%s"
)



func GetStockInfo(stock string, stockType string) (StockInfo, error) {


	// Create and return the StockInfo struct
	stockInfo := StockInfo{
		Stock:     stock,
		Date:      time.Now(),
		Price:     0.0,
		Dividends: []model.Dividend{},
	}


	c := colly.NewCollector()


	fullUrl:= fmt.Sprintf(Investidor10Url, stockType, stock)

	log.Println("Visiting: ", fullUrl)


	c.OnRequest(func(r *colly.Request) { 
		fmt.Println("Visiting: ", r.URL) 
	}) 
	 
	c.OnError(func(_ *colly.Response, err error) { 
		log.Println("Something went wrong: ", err) 
	}) 
	 
	c.OnResponse(func(r *colly.Response) { 
		fmt.Println("Page visited: ", r.Request.URL) 
	}) 

	c.OnHTML("div[class='_card cotacao']", func(e *colly.HTMLElement) {
		e.ForEach("div[class='_card-body']", func(_ int, el *colly.HTMLElement) {			
			price, err :=convtStringToFloat64(el.Text)
			if err != nil {
				log.Println("Erro ao converter preço", err.Error())
			} else {
				stockInfo.Price = price

			}
		},)
	})
	
	c.OnHTML("table[id='table-dividends-history']", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {			
			log.Println("Dividendos: ", el.Text)
			linha:= strings.Split(el.Text, "\n")
			log.Println("Dividendos: ", linha)
			if len(linha) > 6 {
				dateCom, err := time.Parse("02/01/2006", linha[2])
				if err != nil {
					log.Println("Erro ao converter data", err.Error())
				}
				datePay, err := time.Parse("02/01/2006", linha[3])
				if err != nil {
					log.Println("Erro ao converter data", err.Error())
				}
				value, err := convtStringToFloat64(linha[4])
				if err != nil {
					log.Println("Erro ao converter valor", err.Error())
				}
				dividend := model.Dividend{
					Code: stock,
					DateCom: dateCom,
					DatePag: datePay,
					Value: value,
					DividendType: linha[1],
				}
				stockInfo.Dividends = append(stockInfo.Dividends, dividend)
			}
		},)
	})


	c.OnScraped(func(r *colly.Response) { 
		fmt.Println(r.Request.URL, " scraped!") 
	})
	
	c.Visit(fullUrl)	

	return stockInfo, nil
}


func convtStringToFloat64(str string) (float64, error) {
	str = strings.Replace(str, "R$", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, ".", "", -1)
	str = strings.Replace(str, ",", ".", -1)
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Println("Erro ao converter string para float64", err.Error())
		return 0.0, err
	}
	return f, nil
}