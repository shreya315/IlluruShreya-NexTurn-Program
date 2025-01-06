package main

import (
	"errors"
	"fmt"
	"sort"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

var inventory []Product

func addProduct(id int, name string, price interface{}, stock int) error {

	priceFloat, ok := price.(float64)
	if !ok {
		return errors.New("invalid price type; must be a float64")
	}

	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	inventory = append(inventory, Product{ID: id, Name: name, Price: priceFloat, Stock: stock})
	return nil
}

func updateStock(id int, newStock int) error {
	if newStock < 0 {
		return errors.New("stock cannot be negative")
	}

	for i, p := range inventory {
		if p.ID == id {
			inventory[i].Stock = newStock
			return nil
		}
	}
	return errors.New("product not found")
}

func searchProduct(query interface{}) (*Product, error) {
	// Search by ID or name
	for _, p := range inventory {
		switch v := query.(type) {
		case int:
			if p.ID == v {
				return &p, nil
			}
		case string:
			if p.Name == v {
				return &p, nil
			}
		}
	}
	return nil, errors.New("product not found")
}

func displayInventory() {
	fmt.Println("ID\tName\t\tPrice\tStock")
	for _, p := range inventory {
		fmt.Printf("%d\t%s\t\t$%.2f\t%d\n", p.ID, p.Name, p.Price, p.Stock)
	}
}

func sortByPrice() {
	sort.SliceStable(inventory, func(i, j int) bool {
		return inventory[i].Price < inventory[j].Price
	})
}

func sortByStock() {
	sort.SliceStable(inventory, func(i, j int) bool {
		return inventory[i].Stock < inventory[j].Stock
	})
}

func main() {

	err := addProduct(1, "Laptop", 799.99, 10)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = addProduct(2, "Smartphone", 499.49, 20)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = addProduct(3, "Headphones", 150.00, 50)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Initial Inventory:")
	displayInventory()

	err = updateStock(2, 18)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("\nUpdated Inventory (Stock of product 2 updated):")
	displayInventory()

	product, err := searchProduct(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("\nProduct found: %+v\n", *product)
	}

	product, err = searchProduct("Headphones")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("\nProduct found: %+v\n", *product)
	}

	sortByPrice()
	fmt.Println("\nInventory Sorted by Price:")

	sortByStock()
	fmt.Println("\nInventory Sorted by Stock:")
}
