package handles

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HomeHandlerGin(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	var logado bool
	var logedUser string
	if (user==nil){
		logado=false
	} else {
		logado=true
		logedUser=user.(string)
	}

	valor := os.Getenv("goSansoneDB")
	log.Println("HomeHandlerGin...")
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Vari":   "Minha Página com Gin x<"+valor+"> !",
        "Message": "Olá, mundo!",
		"logado" : logado,
		"logedUser" : logedUser,
	})
}

func WelcomeHandlerGin(c *gin.Context){
	c.HTML(http.StatusOK, "welcome.html", nil)
}