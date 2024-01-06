package model

import (
	"log"
)

type Stock struct{
	Code      string `json:"code"`
	Name      string `json:"name"`
	Tipo      string `json:"tipo"`
}


func (stock *Stock) GetStocksList() []Stock{
	db:= GetConnection()
	defer db.Close()

	var stocks []Stock

	rows, err := db.Query("SELECT name, code, type FROM stocks")
	if err != nil {
		log.Fatal(err)
	}
    defer rows.Close()

    for rows.Next() {
        var name string
        var code string
        var tipo string
        err = rows.Scan(&name, &code, &tipo)
        if err != nil {
            log.Fatal(err)
        }
		stocks= append(stocks, Stock{
			Name: name,
			Code: code,
			Tipo: tipo})
    }
	return stocks
}