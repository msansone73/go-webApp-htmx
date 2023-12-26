package handles

import (
	"log"
	"main/model/user"
	"net/http"

	"github.com/gin-contrib/sessions"
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
	session:= sessions.Default(c)
	username := c.PostForm("username")
    password := c.PostForm("password")
	var user user.User
	err := user.GetUserByEmailPass(username, password)
	if err!=nil{
		c.HTML(http.StatusOK,"login.html",nil)
	} else {
		session.Set("user",user.Email)
		err= session.Save()
		if err!= nil{
			log.Println("LoginHandlerPost - impossivel gravar sessão - "+err.Error())
		}
		c.HTML(http.StatusOK,"loginSucesso.html",gin.H{
			"logedUser":username,
		})
	}
}