package application

import (
	"fmt"

	"example.com/m/Users/domain"
)

type RegisterUser struct {
	db domain.IUser
}

func NewRegisterUser(db domain.IUser) *RegisterUser{
	return &RegisterUser{db: db}
}

func(uc *RegisterUser) RegisterUserProcess(username string, email string, password string)error{
	user := domain.User{Username: username, Email: email, Password: password}
	fmt.Println("Si funca")
	return uc.db.SaveUserFunction(user)
}