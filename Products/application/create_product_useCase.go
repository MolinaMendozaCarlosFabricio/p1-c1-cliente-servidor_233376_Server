package aplication

import (
	"fmt"
	"example.com/m/Products/domain"
)

type CreateProduct struct {
	db domain.Iproduct
}

func NewCreateProduct(db domain.Iproduct) *CreateProduct{
	return &CreateProduct{db: db}
}

func (uc *CreateProduct) CreateProductProcess(name string, price float32) error {
	product := &domain.Product{Name: name, Price: price}
	fmt.Println("Si funca")
	return uc.db.Save(product)
}