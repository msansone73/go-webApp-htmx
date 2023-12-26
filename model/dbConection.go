package model

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func getConectionString() string {
	return os.Getenv("goSansoneDB")
}

func getConnection() *sql.DB {
	db, err := sql.Open("postgres", getConectionString())
	if err!=nil {
		log.Fatal(err.Error())
	}

	return db
}

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Stock struct{
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Tipo      string `json:"tipo"`
	
}

func GetUserById(id int) User {
	db:= getConnection()
	defer db.Close()

	var user User
	err := db.QueryRow("SELECT id, name, email, password FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Fatal(err)
	}

	return user

}

func GetStocksList() []Stock{
	db:= getConnection()
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