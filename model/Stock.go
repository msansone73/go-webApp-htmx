package model

import (
	"log"
)

type Stock struct{
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Tipo      string `json:"tipo"`
}


func (stock *Stock) GetStocksList() []Stock{
	db:= GetConnection()
	defer db.Close()

	var stocks []Stock

	rows, err := db.Query("SELECT id, name, code, type FROM stocks")
	if err != nil {
		log.Fatal(err)
	}
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        var code string
        var tipo string
        err = rows.Scan(&id, &name, &code, &tipo)
        if err != nil {
            log.Fatal(err)
        }
		stocks= append(stocks, Stock{
			ID: id,
			Name: name,
			Code: code,
			Tipo: tipo})
    }
	return stocks
}