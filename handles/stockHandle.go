package handles

import (
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetStocks( c *gin.Context) {
	stock := new(model.Stock)
	stocks := stock.GetStocksList()
	c.HTML(http.StatusOK,"stockList.html", stocks)

}

