package aplication

import (
	"fmt"

	"example.com/m/Products/domain"
)

type GetLastAddedProduct struct {
	db domain.Iproduct
}

func NewGetLastAddedProduct(db domain.Iproduct)*GetLastAddedProduct{
	return&GetLastAddedProduct{db: db}
}

func(uc *GetLastAddedProduct)GetLastAddedProductProcess()(domain.Product, error){
	fmt.Println("Si funca")
	result, err := uc.db.GetLastAddedProduct()
	return result, err
}