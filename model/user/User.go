package user

import (
	"log"
	"main/model"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (user *User) GetUserById(id int) error {
	
	db:= model.GetConnection()
	defer db.Close()

	err := db.QueryRow("SELECT id, name, email,  FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		log.Println("GetUserById() - "+err.Error())
		return err
	}
	return nil
}

func (user *User) GetUserByEmail(email string) error {
	
	db:= model.GetConnection()
	defer db.Close()

	err := db.QueryRow("SELECT id, name, email FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		log.Println("GetUserById() - "+err.Error())
		return err
	}
	return nil
}

func (user *User) GetUserByEmailPass(email string, password string) error {
	
	db:= model.GetConnection()
	defer db.Close()

	err := db.QueryRow("SELECT id, name, email FROM users WHERE email = $1 and password= $2", email, password).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		log.Println("GetUserById() - "+err.Error())
		return err
	}
	return nil
}


