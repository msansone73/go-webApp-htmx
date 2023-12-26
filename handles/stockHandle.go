package handles

import (
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetStocks( c *gin.Context) {
	stocks := model.GetStocksList()
	c.HTML(http.StatusOK,"stockList.html", stocks)

}

