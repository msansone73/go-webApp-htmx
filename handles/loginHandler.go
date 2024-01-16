package handles

import (
	"log"
	"main/model"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func LoginSucessoHandle(c *gin.Context) {
	logedUser, _:= c.Cookie("Value")
	c.HTML(http.StatusOK, "loginSucesso.html", gin.H{
		"logedUser":logedUser,
	})
}

func LogoutHandler(c * gin.Context){
	session:= sessions.Default(c)
	session.Clear()
	session.Save()
	log.Println("LogoutHandler - Session cleaned!")
	c.HTML(http.StatusOK, "forbidden.html", gin.H{
		"mensagem":"logout efetuado com sucesso!",
	})
}

func LoginHandlerGet(c *gin.Context){
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginHandlerPost(c *gin.Context){
	session:= sessions.Default(c)
	username := c.PostForm("username")
    password := c.PostForm("password")
	var user model.User
	err := user.GetUserByEmailPass(username, password)
	if err!=nil{
		c.HTML(http.StatusOK,"login.html",nil)
	} else {
		session.Clear()
		session.Options(sessions.Options{
			MaxAge: 3600 * 1, // 1 hours
			Path:   "/",
			Secure: true,
		})
		session.Set("user",user.Email)
		log.Println("LoginHandlerPost - user: "+user.Name)
		log.Println("LoginHandlerPost - ID: "+strconv.Itoa(user.ID))
		err= session.Save()
		if err!= nil{
			log.Println("LoginHandlerPost - impossivel gravar sessão - "+err.Error())
		}
		c.HTML(http.StatusOK,"loginSucesso.html",gin.H{
			"logedUser":username,
		})
	}
}