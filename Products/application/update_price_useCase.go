package aplication

import (
	"fmt"

	"example.com/m/Products/domain"
)

type UpdatePrice struct {
	db domain.Iproduct
}

func NewUpdatePrice(db domain.Iproduct)*UpdatePrice{
	return&UpdatePrice{db: db}
}

func(uc *UpdatePrice)UpdatePriceProcess(id int32, price float32) error{
	fmt.Println("Si funca")
	err := uc.db.UpdatePriceProduct(id, price)
	return err
}