package main

import (
	"fmt"
	"main/model"
	"main/scrapping"
)

func main() {
	fmt.Println("Starting the application...")
	stock := model.Stock{}
	stocks:= stock.GetStocksList()

	for _, stock := range stocks {
		fmt.Println(stock)
		typeOfStock := "acoes"
		if stock.Tipo == "FII" {
			typeOfStock = "fii"
		}
		stockInfo, err := scrapping.GetStockInfo(stock.Code, typeOfStock)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			fmt.Println("saving:",stockInfo.Stock)
			dividend := new(model.Dividend)
			dividend.SaveDividend(stockInfo.Dividends)
			fmt.Println("saved:",stockInfo.Stock)

		}
		
	}
	fmt.Println("Terminating the application...")
}	
