package handles

import (
	"html/template"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := PageData{
        Vari:   "Minha Página",
        Message: "Olá, mundo!",
    }
	tmpl.Execute(w, data)
}

func HomeHandlerGin(c *gin.Context) {
	log.Println("HomeHandlerGin...")
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Vari":   "Minha Página com Gin",
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
	c.HTML(http.StatusOK, "about.html", gin.H{
        "Usuario":   sessionToken,
        "Message": "Olá, mundo!",
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