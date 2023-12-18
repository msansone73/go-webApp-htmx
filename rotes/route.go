package rotes

import (
	"main/handles"
	"github.com/gin-gonic/gin"
)




func SetRoutesGin(r *gin.Engine ){

	r.GET("/", handles.HomeHandlerGin)
	r.GET("/about", handles.AboutHandler)
	r.GET("/loginSucesso", handles.LoginSucessoHandle)
	r.GET("/login", handles.LoginHandlerGet)
	r.POST("/login", handles.LoginHandlerPost)

}