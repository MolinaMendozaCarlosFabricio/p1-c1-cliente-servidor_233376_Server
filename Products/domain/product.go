package domain

type Product struct {
	ID int32
	Name string
	Price float32
}

func NewProduct(name string, price float32) *Product {
	return &Product{ID: 1, Name: name, Price: price}
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) SetName(name string){
	p.Name = name
}