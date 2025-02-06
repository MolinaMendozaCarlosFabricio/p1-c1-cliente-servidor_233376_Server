package aplication

import (
	"fmt"

	"example.com/m/Products/domain"
)

type EditProduct struct {
	db domain.Iproduct
}

func NewEditProduct(db domain.Iproduct) *EditProduct{
	return &EditProduct{db:db}
}

func(uc *EditProduct) EditProductProcess(name string, price float32, id int32)error{
	fmt.Println("Si funca")
	return uc.db.Edit(name, price, id)
}