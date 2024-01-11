package main

import (
	"fmt"
	"html/template"
	"log"
	"main/rotes"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)


func main() {


	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
        MaxAge: 3600 * 1, // 1 hours
        Path:   "/",
        Secure: true,
    })
    // Configurar o Gin para usar o logger padrão
    //gin.SetMode(gin.ReleaseMode)
    //gin.DefaultWriter = file


	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"formatCurrency": formatCurrency,
		"formatDate": formatDate,
	})
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

func formatCurrency(input float64) string {
    // Format the number as a string with 2 decimal places
    s := fmt.Sprintf("%.2f", input)

    // Split the string into the integer part and the decimal part
    parts := strings.Split(s, ".")

    // Add commas as thousand separators to the integer part
    for i := len(parts[0]) - 3; i > 0; i -= 3 {
        parts[0] = parts[0][:i] + "." + parts[0][i:]
    }

    // Join the parts back together
    return "$" + parts[0] + "," + parts[1]
}


// create a function that format a date like dd/mm/yyyy
func formatDate(t time.Time) string {


	return t.Format("02/01/2006")
}