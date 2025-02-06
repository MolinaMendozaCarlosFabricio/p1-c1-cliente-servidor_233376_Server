package application

import (
	"fmt"

	"example.com/m/Users/domain"
)

type EditUser struct {
	db domain.IUser
}

func NewEditUser(db domain.IUser)*EditUser{
	return &EditUser{db: db}
}

func(uc *EditUser) EditUserProcess(id int32, username string, email string, password string)error{
	fmt.Println("Si funca")
	user := &domain.User{ID: id, Username: username, Email: email, Password: password}
	err := uc.db.EditUserFunction(*user)
	return err
}