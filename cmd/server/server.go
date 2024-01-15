package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"main/rotes"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)


func init() {
	setLog()
}

func main() {
	r := gin.Default()
	setFormaters(r)
	setSession(r)
	rotes.SetRoutesGin(r)
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./assets")
	r.Run(":80")
}

func setFormaters(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		"formatCurrency": formatCurrency,
		"formatDate": formatDate,
	})
}

func setSession(r *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
        MaxAge: 3600 * 1, // 1 hours
        Path:   "/",
        Secure: true,
    })
	r.Use(sessions.Sessions("mysession", store))

}

func setLog(){
	logfile := "app.log"
	f, err := os.Create(logfile)
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f)

    log.SetFlags(log.LstdFlags | log.Lshortfile)
    file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal("Falha ao abrir arquivo de log:", err)
    }
    log.SetOutput(file)
}

func formatCurrency(input float64) string {
    s := fmt.Sprintf("%.2f", input)
    parts := strings.Split(s, ".")
    for i := len(parts[0]) - 3; i > 0; i -= 3 {
        parts[0] = parts[0][:i] + "." + parts[0][i:]
    }
    return "$" + parts[0] + "," + parts[1]
}

func formatDate(t time.Time) string {
	return t.Format("02/01/2006")
}