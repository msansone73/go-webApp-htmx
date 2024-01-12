package handles

import (
	"log"
	"main/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type TransactionFormReq struct {
	Stocks      []model.Stock     `json:"stocks"`
	Transaction model.Transaction `json:"transaction"`
}

func TransactionListHandler(c *gin.Context) {
	log.Println("iniciando TransactionListHandler() ...")
	session := sessions.Default(c)
	iemail := session.Get("user")
	email := iemail.(string)
	var user model.User
	err := user.GetUserByEmail(email)
	if err != nil {
		log.Println("GetTransactions - Erro ao recuperar email - email=" + email)
	}
	tran := new(model.Transaction)
	trans, err := tran.GetTransactionsByUser(user)
	if err != nil {
		log.Println("GetTransactions - Erro ao recuperar email - email=" + email)
	}
	c.HTML(http.StatusOK, "transactionList.html", trans)
}

type TransactionItem struct {
	Code  string  `json:"code"`
	Valor float64 `json:"valor"`
}

type Carteira struct {
	Transactions []TransactionItem `json:"transactions"`
	ValorTotal   float64           `json:"valorTotal"`
}

func CarteiraHandler(c *gin.Context) {
	log.Println("iniciando CarteiraHandler() ...")
	session := sessions.Default(c)
	iemail := session.Get("user")
	email := iemail.(string)
	var user model.User
	err := user.GetUserByEmail(email)
	if err != nil {
		log.Println("GetTransactions - Erro ao recuperar email - email=" + email)
	}
	tran := new(model.Transaction)
	trans, err := tran.GetTransactionsByUser(user)
	if err != nil {
		log.Println("GetTransactions - Erro ao recuperar email - email=" + email)
	}

	transactionItens := map[string]float64{}
	for _, t := range trans {
		if transactionItens[t.Stock.Code] == 0 {
			transactionItens[t.Stock.Code] = t.Value * float64(t.Quantity)
		} else {
			transactionItens[t.Stock.Code] += t.Value * float64(t.Quantity)
		}
	}
	c.HTML(http.StatusOK, "carteira.html", transactionItens)
}

func TransactionFormHandler(c *gin.Context) {
	log.Println("iniciando TransactionFormHandler() ...")
	id, _ := strconv.Atoi(c.Param("id"))

	tranReq := new(TransactionFormReq)
	if id != 0 {
		tranReq.Transaction.GetTransactionById(id)
	} else {
		tranReq.Transaction.Date = time.Now()
	}
	stock := new(model.Stock)
	stocks := stock.GetStocksList()
	tranReq.Stocks = stocks

	c.HTML(http.StatusOK, "transactionForm.html", tranReq)
}

func TransactionFormHandlerPost(c *gin.Context) {
	log.Println("iniciando TransactionFormHandlerPost() ...")
	transaction := new(model.Transaction)
	transaction.ID, _ = strconv.Atoi(c.PostForm("id"))
	transaction.Stock.Code = c.PostForm("stock")
	transaction.Quantity, _ = strconv.Atoi(c.PostForm("quantity"))
	transaction.Value, _ = strconv.ParseFloat(c.PostForm("value"), 64)
	transaction.Tipo = c.PostForm("type")
	dateStr := c.PostForm("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Println("Error parsing date:", err)
		// Handle the error accordingly
	}
	transaction.Date = date

	if transaction.ID != 0 {
		//transaction.UpdateTransaction()
	} else {
		session := sessions.Default(c)
		iemail := session.Get("user")
		email := iemail.(string)
		var user model.User
		err := user.GetUserByEmail(email)
		if err != nil {
			log.Println("GetTransactions - Erro ao recuperar email - email=" + email)
		}
		transaction.User = user
	}
	transaction.SaveTransaction()
	c.Redirect(http.StatusFound, "/transactions")
}

func TransactionDeleteHandler(c *gin.Context) {
	log.Println("iniciando TransactionDeleteHandler() ...")
	id, _ := strconv.Atoi(c.Param("id"))
	transaction := new(model.Transaction)
	transaction.ID = id
	transaction.DeleteTransaction()
	TransactionListHandler(c)
}