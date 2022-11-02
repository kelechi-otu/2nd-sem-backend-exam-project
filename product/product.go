package product

import (
	"errors"
	"fmt"
)

type car struct {
	model  string
	color string
}

func (c car) Model() string {
	return c.model
}

type Product interface {
	Price() float32
	Quantity() int
	Name() string
	IsAvaliable() bool
	Display()
	Sell(quanity int) error
}

type product struct {
	car
	name     string
	quantity int
	price    float32
}

func New(carModel string, color string, name string, quantity int, price float32) *product {
	return &product{
		car: car{
			model: carModel,
			color: color,
		},
		name:     name,
		quantity: quantity,
		price:    price,
	}
}

func (p product) Price() float32 {
	return p.price
}

func (p product) Quantity() int {
	return p.quantity
}

func (p product) Name() string {
	return p.name
}

func (p product) IsAvaliable() bool {
	if p.quantity == 0 {
		return false
	}
	return true
}

func (p product) Display() {
	fmt.Printf("product: %v", p)
}

func (p *product) Sell(quantity int) error {
	if p.quantity < quantity {
		return errors.New("not enough products to sell")
	}

	p.quantity = p.quantity - quantity

	return nil
}
