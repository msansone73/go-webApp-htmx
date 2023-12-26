package handles

import (
	"log"
	"main/model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)


func AboutHandler(c *gin.Context) {
	var sessionToken string
	var user = new(user.User)
	
	cookie, err := c.Cookie("Value")
	if err != nil {
		sessionToken="não logado"
	} else {
		sessionToken = cookie
	}

	err = user.GetUserById(10)
	if err!=nil {
		log.Println("AboutHandler() - "+err.Error())
	}

	c.HTML(http.StatusOK, "about.html", gin.H{
        "Usuario":   sessionToken,
        "Message": "Olá, mundo!",
		"user":user,
	})
}