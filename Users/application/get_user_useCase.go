package application

import (
	"fmt"

	"example.com/m/Users/domain"
)

type GetUser struct {
	db domain.IUser
}

func NewGetUser(db domain.IUser) *GetUser{
	return &GetUser{db: db}
}

func (uc *GetUser) GetUserProcess(id int32)([]domain.User, error){
	fmt.Println("Si funca")
	result, err := uc.db.GetUserFunction(id)
	return result, err	
}