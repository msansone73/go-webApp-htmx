package handles

import (
	"log"
	"main/model"
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

func AboutHandler(c *gin.Context) {
	var sessionToken string
	
	cookie, err := c.Cookie("Value")
	if err != nil {
		sessionToken="não logado"
	} else {
		sessionToken = cookie
	}

	user := model.GetUserById(1)

	c.HTML(http.StatusOK, "about.html", gin.H{
        "Usuario":   sessionToken,
        "Message": "Olá, mundo!",
		"user":user,
	})
}

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
	username := c.PostForm("username")
    password := c.PostForm("password")
	if password!="" {
		c.SetCookie("Value",username,0,"","",true,true)
		c.HTML(http.StatusOK,"loginSucesso.html",gin.H{
			"logedUser":username,
		})
	} else {
		c.HTML(http.StatusOK,"login.html",nil)
	}
}

func GetStocks( c *gin.Context) {
	stocks := model.GetStocksList()
	c.HTML(http.StatusOK,"stockList.html", stocks)

}