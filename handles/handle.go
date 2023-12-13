package handles

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)



func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := PageData{
        Vari:   "Minha Página",
        Message: "Olá, mundo!",
    }
	tmpl.Execute(w, data)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	var sessionToken string
	cookie, err := r.Cookie("session_token")
	if err != nil {
		sessionToken="não logado"
	} else {
		sessionToken = cookie.Value
	}

	data := PageData{
        Usuario:   sessionToken,
        Message: "Olá, mundo!",
    }

	tmpl := template.Must(template.ParseFiles("templates/about.html"))
	tmpl.Execute(w, data)
}

func LoginSucessoHandle(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/loginSucesso.html"))
	tmpl.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	switch r.Method {
    case "GET":
        // Manipulação para requisições GET
		templ:= template.Must(template.ParseFiles("templates/login.html"))
		templ.Execute(w,nil)
    case "POST":
        // Manipulação para requisições POST
        // Ler o corpo da requisição
        body, err := io.ReadAll(r.Body)
		if err != nil {
            http.Error(w, "Erro ao ler o corpo da requisição", http.StatusInternalServerError)
            return
        }
		logger.Println(string(body))
		valores, err := url.ParseQuery(string(body))
		if err != nil {
            http.Error(w, "Erro ao parsear o corpo da requisição", http.StatusInternalServerError)
            return
        }
		username := valores.Get("username")
    	password := valores.Get("password")
        if username=="marcio" && password=="" {

			http.SetCookie(w, &http.Cookie{
				Name: "session_token",
				Value: "Marcio",
			})

			templ:= template.Must(template.ParseFiles("templates/loginSucesso.html"))
			templ.Execute(w,nil)
		} else {
			templ:= template.Must(template.ParseFiles("templates/login.html"))
			templ.Execute(w,nil)
		}
    default:
        // Método não suportado
        http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
    }


}