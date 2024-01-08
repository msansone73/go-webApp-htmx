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

	r.GET("/transactions", handles.TransactionListHandler)

	r.GET("/carteira", handles.CarteiraHandler)

}