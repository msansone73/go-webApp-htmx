package main

import (
	"main/rotes"
    "github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	rotes.SetRoutesGin(r)
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./assets")
	
	r.Run(":80")


}






