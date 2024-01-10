package handles

import (
	"log"
	"main/model"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type TransactionFormReq struct {
	Stocks []model.Stock `json:"stocks"`
	Transaction model.Transaction `json:"transaction"`
}

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

type TransactionItem struct {
	Code string `json:"code"`
	Valor float64 `json:"valor"`
}

type Carteira struct {
	Transactions []TransactionItem `json:"transactions"`
	ValorTotal float64 `json:"valorTotal"`
}

func CarteiraHandler (c *gin.Context){
	log.Println("iniciando CarteiraHandler() ...")
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

	transactionItens := map[string]float64{}
	for _, t:= range trans {
		if transactionItens[t.Stock.Code] == 0 {
			transactionItens[t.Stock.Code] = t.Value*float64(t.Quantity)
		} else {
			transactionItens[t.Stock.Code] += t.Value*float64(t.Quantity)
		}	
	}
	c.HTML(http.StatusOK,"carteira.html", transactionItens)
}

func TransactionFormHandler(c *gin.Context) {
	log.Println("iniciando TransactionFormHandler() ...")
	id, _ := strconv.Atoi(c.Param("id"))

	tranReq := new(TransactionFormReq)
	if id!= 0 {
		tranReq.Transaction.GetTransactionById(id)
	}
	stock := new(model.Stock)
	stocks := stock.GetStocksList()
	tranReq.Stocks = stocks

	c.HTML(http.StatusOK, "transactionForm.html", tranReq)
}