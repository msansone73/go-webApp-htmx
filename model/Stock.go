package model

import (
	"log"
)

type Stock struct{
	Code      string `json:"code"`
	Name      string `json:"name"`
	Tipo      string `json:"tipo"`
}

func (stock *Stock) GetStocksList() []Stock {
	db := GetConnection()
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
		stocks = append(stocks, Stock{
			Name: name,
			Code: code,
			Tipo: tipo,
		})
	}
	return stocks
}


func (stock *Stock) SaveStock() error {
	db := GetConnection()
	defer db.Close()

	// Check if the stock code already exists in the database
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM stocks WHERE code = $1", stock.Code).Scan(&count)
	if err != nil {
		log.Printf("Error checking if stock code exists: %s", err)
		return err
	}
	if count > 0 {
		// Update the existing stock
		_, err = db.Exec("UPDATE stocks SET name = $1, type = $2 WHERE code = $3", stock.Name, stock.Tipo, stock.Code)
		if err != nil {
			log.Printf("Error updating stock: %s", err)
			return err
		}
	} else {
		// Insert a new stock
		_, err = db.Exec("INSERT INTO stocks (name, code, type) VALUES ($1, $2, $3)", stock.Name, stock.Code, stock.Tipo )
		if err != nil {
			log.Printf("Error inserting stock: %s", err)
			return err
		}
	}

	return nil
}


func (stock *Stock) DeleteStock() error {
	db := GetConnection()
	defer db.Close()

	_, err := db.Exec("DELETE FROM stocks WHERE code = $1", stock.Code)
	if err != nil {
		log.Printf("Error deleting stock: %s", err)
		return err
	}

	return nil
}

func (stock *Stock) GetStockByCode() error {
	db := GetConnection()
	defer db.Close()

	err := db.QueryRow("SELECT name, code, type FROM stocks WHERE code = $1", stock.Code).Scan(&stock.Name, &stock.Code, &stock.Tipo)
	if err != nil {
		log.Printf("Error getting stock: %s", err)
		return err
	}

	return nil
}