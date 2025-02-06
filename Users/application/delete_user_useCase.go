package application

import (
	"fmt"

	"example.com/m/Users/domain"
)

type DeleteUser struct {
	db domain.IUser
}

func NewDeleteUser(db domain.IUser)*DeleteUser{
	return &DeleteUser{db: db}
}

func(uc *DeleteUser) DeleteUserProcess(id int32)error{
	fmt.Println("Si funca")
	err := uc.db.DeleteUserFunction(id)
	return err
}