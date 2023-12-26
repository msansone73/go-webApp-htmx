package handles

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HomeHandlerGin(c *gin.Context) {

	valor := os.Getenv("goSansoneDB")
	log.Println("HomeHandlerGin...")
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Vari":   "Minha Página com Gin x<"+valor+"> !",
        "Message": "Olá, mundo!",
	})
}