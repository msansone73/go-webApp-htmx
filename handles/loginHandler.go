package handles

import (
	"main/model/user"
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
	var user user.User
	err := user.GetUserByEmailPass(username, password)
	if err!=nil{
		c.HTML(http.StatusOK,"login.html",nil)
	} else {
		c.SetCookie("Value",username,0,"","",true,true)
		c.HTML(http.StatusOK,"loginSucesso.html",gin.H{
			"logedUser":username,
		})
	}
}