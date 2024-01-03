package model

import (
	"log"
	"time"
)

type Transaction struct {
	ID			int 	`json:"id"`
	User 		User 	`json:"User"`
	Stock 		Stock	`json:"Stock"`
	Tipo		string	`json:"type"`
	Value		float64	`json:"value"`
	Quantity	int		`json:"quantity"`
	Date		time.Time `json:"data"`
}

func (t * Transaction) GetTransactionsByUser(user User) ([]Transaction, error) {
	db:= GetConnection()
	defer db.Close()


	rows, err:= db.Query("select id, stock_id, type, value, quantity, data_at from transactions where user_id = $1", user.ID) 
	if err!= nil {
		log.Println("GetTransactionsByUser - erro: "+err.Error())
		return nil, err
	}

	var transactions = make([]Transaction,0)

	for rows.Next() {
		var id, quantity, stock_id int
		var tipo string
		var value float64
		var data time.Time

		err = rows.Scan(&id, &stock_id, &tipo, &value, &quantity, &data)
		if err!= nil {
			log.Println("GetTransactionsByUser - erro: "+err.Error())
			return nil, err
		}

		transactions = append(transactions, Transaction{
			ID: id,
			User: user,
			Stock: Stock{
				ID: stock_id,
				},
			Tipo: tipo,
			Value: value,
			Quantity: quantity,
			Date: data,
		})
	}
	return transactions, nil
}