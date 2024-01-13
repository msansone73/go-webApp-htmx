package rotes

import (
	"main/handles"
	"github.com/gin-gonic/gin"
)

func SetRoutesGin(r *gin.Engine ){

	r.GET("/", handles.HomeHandlerGin)
	r.GET("/welcome", handles.MustLoggged, handles.WelcomeHandlerGin)

	r.GET("/about", handles.AboutHandler)

	r.GET("/loginSucesso", handles.LoginSucessoHandle)
	r.GET("/login", handles.LoginHandlerGet)
	r.POST("/login", handles.LoginHandlerPost)
	r.GET("/logout", handles.LogoutHandler)
	r.GET("/forbidden", handles.GoForbidden)

	r.GET("/stocks", handles.MustLoggged, handles.GetStocks)
	r.GET("/stock/form", handles.MustLoggged, handles.GetStocksForm)
	r.POST("/stock/form", handles.MustLoggged, handles.PostStocksForm)
	r.DELETE("/stock/:code", handles.MustLoggged, handles.DeleteStock)
	r.PUT("/stock/:code", handles.MustLoggged, handles.PutStock)
	r.GET("/stock/scraping/:code", handles.MustLoggged, handles.StocksScrapping)

	r.GET("/transactions", handles.MustLoggged, handles.TransactionListHandler)
	r.GET("/transaction/form", handles.MustLoggged, handles.TransactionFormHandler)
	r.GET("/transaction/:id", handles.MustLoggged, handles.TransactionFormHandler)
	r.POST("/transaction/form", handles.MustLoggged, handles.TransactionFormHandlerPost)
	r.DELETE("/transaction/:id", handles.MustLoggged, handles.TransactionDeleteHandler)

	r.GET("/carteira", handles.MustLoggged, handles.CarteiraHandler)

}