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

	err := db.QueryRow("SELECT id, name, email, password FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println("GetUserById() - "+err.Error())
		return err
	}
	return nil
}


