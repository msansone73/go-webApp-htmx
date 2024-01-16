package handles

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func GoForbidden(c *gin.Context){
	log.Println("GoForbidden - sem acesso - IP=" + c.ClientIP())
	c.HTML(http.StatusOK, "forbidden.html", gin.H{
		"mensagem":"sem acesso",
	})
}

func MustLoggged( c *gin.Context ){
	session:= sessions.Default(c)
	email := session.Get("user")
	log.Println("MustLoggged - email: "+email.(string))
	if (email==nil){
		log.Println("MustLoggged - Usuario Nâo logado!")
		c.Redirect( 302,"/forbidden")
	} else {
		c.Next()
	}
}