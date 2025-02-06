package aplication

import (
	"fmt"

	"example.com/m/Products/domain"
)

type GetAllProducts struct {
	db domain.Iproduct
}

func NewGetAllProducts(db domain.Iproduct) *GetAllProducts{
	return &GetAllProducts{db: db}
}

func (uc *GetAllProducts) GetAllProductProcess() ([]domain.Product, error) {
	fmt.Println("Si funca")
	return uc.db.GetAll()
}