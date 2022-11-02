package main

import (
	"backend-project/product"
	"backend-project/store"
	"fmt"
)

func main() {
	// create a new product 
	prod1 := product.New("honda", "red", "product1", 10, 5000)
	prod2 := product.New("hilux", "grey", "product2", 50, 8000)

	products := []product.Product{prod1, prod2}

	// create a new store
	store := store.New("random-store", products)

	// Add new items to the store
	prod3 := product.New("venza", "black", "product3", 95, 34000)
	store.AddItemsToStore(prod3)

	// sell items
	err := store.SellProduct("venza", 7)
	if err != nil {
		fmt.Println(err)
	}

	// He needs to see number of cars left to be sold
	store.NumberOfCarsLeftToBeSold()

	// He needs to see the sum of prices of cars left
	store.SumOfPricesOfCarsLeft()

	// sum total of prices of cars sold
	store.TotalPriceOfCarsSold()

	// number of cars he has sold
	store.NumbersOfCarsSold()

	// list of orders of sales made
	store.ListOfCarsSold()

}