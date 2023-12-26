package handles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func LoginSucessoHandle(c *gin.Context) {
	logedUser, _:= c.Cookie("Value")
	c.HTML(http.StatusOK, "loginSucesso.html", gin.H{
		"logedUser":logedUser,
	})
}

func LoginHandlerGet(c *gin.Context){
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginHandlerPost(c *gin.Context){
	username := c.PostForm("username")
    password := c.PostForm("password")
	if password!="" {
		c.SetCookie("Value",username,0,"","",true,true)
		c.HTML(http.StatusOK,"loginSucesso.html",gin.H{
			"logedUser":username,
		})
	} else {
		c.HTML(http.StatusOK,"login.html",nil)
	}
}