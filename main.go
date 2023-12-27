package main

import (
	"log"
	"main/rotes"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var file *os.File

func main() {

	store := cookie.NewStore([]byte("secret"))

    // Configurar o Gin para usar o logger padrão
    gin.SetMode(gin.ReleaseMode)
    gin.DefaultWriter = file


	r := gin.Default()
	r.Use(sessions.Sessions("mysession", store))

	rotes.SetRoutesGin(r)
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./assets")
	
	r.Run(":80")
}

func init() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal("Falha ao abrir arquivo de log:", err)
    }
    log.SetOutput(file)
}





