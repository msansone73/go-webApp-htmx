package handles

import (
	"main/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func AboutHandler(c *gin.Context) {
	var session = sessions.Default(c)
	var user = new(model.User)
	
	email := session.Get("user")
	if (email==nil){
		user.Email="-- not logged --"
	} else {
		user.Email=email.(string)
	}

	c.HTML(http.StatusOK, "about.html", gin.H{
        "Usuario":   user.Email,
        "Message": "Olá, mundo!",
		"user": user.Email,
	})
}