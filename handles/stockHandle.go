package handles

import (
	"log"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetStocks( c *gin.Context) {
	stock := new(model.Stock)
	stocks := stock.GetStocksList()
	c.HTML(http.StatusOK,"stockList.html", stocks)

}

func GetStocksForm( c *gin.Context) {
	c.HTML(http.StatusOK,"stockForm.html", nil)
}

func PostStocksForm( c *gin.Context) {
	stock := new(model.Stock)
	stock.Name = c.PostForm("name")
	stock.Code = c.PostForm("code")
	stock.Tipo = c.PostForm("tipo")
	err := stock.SaveStock()
	if err!=nil {
		log.Println("Erro ao salvar stock", err.Error())
		c.HTML(http.StatusOK,"stockForm.html", gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/stocks")
}

func DeleteStock( c *gin.Context) {
	code := c.Param("code")
	stock := new(model.Stock)
	stock.Code = code
	err := stock.DeleteStock()
	if err!=nil {
		log.Println("Erro ao deletar stock", err.Error())
	}
	GetStocks(c)
}

func PutStock( c *gin.Context) {
	code := c.Param("code")
	stock := new(model.Stock)
	stock.Code = code
	err := stock.GetStockByCode()
	if err!=nil {
		log.Println("Erro ao recuperar stock", err.Error())
	}
	c.HTML(http.StatusOK,"stockForm.html", stock)
}