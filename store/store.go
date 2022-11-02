package store

import (
	"backend-project/product"
	"errors"
)

type Store interface {
	NumberOfCarsLeftToBeSold() int
	SumOfPricesOfCarsLeft() float32
	NumbersOfCarsSold() int
	TotalPriceOfCarsSold() float32
	ListOfCarsSold() []product.Product
	AddItemsToStore(product product.Product)
	SellProduct(productName string, quantity int) error
}

type store struct {
	name         string
	products     []product.Product
	productsSold []product.Product
}

func New(name string, products []product.Product) *store {
	return &store{
		name:         name,
		products:     products,
		productsSold: []product.Product{},
	}
}

func (s store) NumberOfCarsLeftToBeSold() int {
	var totalQuantity int
	for _, prod := range s.products {
		totalQuantity += prod.Quantity()
	}
	return totalQuantity
}

func (s store) SumOfPricesOfCarsLeft() float32 {
	var totalPrice float32
	for _, prod := range s.products {
		totalPrice += prod.Price() * float32(prod.Quantity())
	}
	return totalPrice
}

func (s store) NumbersOfCarsSold() int {
	var totalNumber int
	for _, prod := range s.productsSold {
		totalNumber += prod.Quantity()
	}
	return totalNumber
}

func (s store) TotalPriceOfCarsSold() float32 {
	var totalPrice float32
	for _, prod := range s.productsSold {
		totalPrice += prod.Price()
	}
	return totalPrice
}

func (s store) ListOfCarsSold() []product.Product {
	return s.productsSold
}

func (s store) AddItemsToStore(product product.Product) {
	s.products = append(s.products, product)
}

func (s store) SellProduct(productName string, quantity int) error {
	var product product.Product
	for _, p := range s.products {
		if p.Name() == productName {
			product = p
		}
	}

	if product == nil {
		return errors.New("product id not found")
	}

	err := product.Sell(quantity)
	if product == nil {
		return err
	}

	//update products sold
	s.productsSold = append(s.productsSold, product)

	return nil
}
