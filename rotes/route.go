package rotes

import (
	"main/handles"
	"net/http"
)


func SetRotes(htt *http.ServeMux ){
	htt.HandleFunc("/", handles.HomeHandler)
	htt.HandleFunc("/about", handles.AboutHandler)
	htt.HandleFunc("/login", handles.LoginHandler)
	htt.HandleFunc("/loginSucesso", handles.LoginSucessoHandle)
}