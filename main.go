package main

import (
	"main/rotes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	store := cookie.NewStore([]byte("secret"))

	r := gin.Default()
	r.Use(sessions.Sessions("mysession", store))

	rotes.SetRoutesGin(r)
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./assets")
	
	r.Run(":80")


}






