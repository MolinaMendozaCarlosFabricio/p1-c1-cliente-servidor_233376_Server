package application

import (
	"fmt"

	"example.com/m/Users/domain"
)

type ComprobateUsers struct {
	db domain.IUser
}

func NewComprobateUsers(db domain.IUser)*ComprobateUsers{
	return &ComprobateUsers{db: db}
}

func(uc *ComprobateUsers)ComprobateUsersProcess()(int, error){
	fmt.Println("Si funca")
	quantity_of_users, err := uc.db.ComprobateQuantityOfUsers()
	return quantity_of_users, err
}