package aplication

import (
	"fmt"

	"example.com/m/Products/domain"
)

type DeleteProduct struct {
	db domain.Iproduct
}

func NewDeleteProduct(db domain.Iproduct)*DeleteProduct{
	return &DeleteProduct{db: db}
}

func(r *DeleteProduct)DeleteProductProcess(id int32)error{
	fmt.Println("Si funca")
	return r.db.Delete(id)
}