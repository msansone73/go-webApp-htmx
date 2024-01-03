package handles

import (
	"log"
	"main/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func TransactionListHandler( c *gin.Context) {
	log.Println("iniciando TransactionListHandler() ...")
	session:= sessions.Default(c)
	iemail := session.Get("user")
	email:= iemail.(string)
	var user model.User
	err := user.GetUserByEmail(email)
	if err!=nil {
		log.Println("GetTransactions - Erro ao recuperar email - email="+email)
	}

	tran := new(model.Transaction)
	trans, err := tran.GetTransactionsByUser(user)
	if err!=nil {
		log.Println("GetTransactions - Erro ao recuperar email - email="+email)
	}
	c.HTML(http.StatusOK,"transactionList.html", trans)

}

func CarteiraHandler (c *gin.Context){


	c.HTML(http.StatusOK,"carteira.html", nil)
}