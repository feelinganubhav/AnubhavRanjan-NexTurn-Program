/*
Exercise 3: Inventory Management System
Topics Covered: Go Conditions, Go Type Casting, Go Functions, Go Arrays, Go Strings,
Go Errors
Case Study:
A store needs to manage its inventory of products. Build an application that includes
the following:
1. Product Struct: Create a struct to represent a product with fields for ID, name,
price (float64), and stock (int).
2. Add Product: Write a function to add new products to the inventory. Use type
casting to ensure price inputs are converted to float64.
3. Update Stock: Implement a function to update the stock of a product. Use
conditions to validate the input (e.g., stock cannot be negative).
4. Search Product: Allow users to search for products by name or ID. If a product is
not found, return a custom error message.
5. Display Inventory: Use loops to display all available products in a formatted
table.
Bonus:
• Add sorting functionality to display products by price or stock in ascending order.
*/

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

func addProduct(id int, name string, price float64, stock int) error {
	if id <= 0 {
		return errors.New("product ID must be positive")
	}
	if price < 0 {
		return errors.New("price cannot be negative")
	}
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}
	for _, product := range inventory {
		if product.ID == id {
			return errors.New("product ID must be unique")
		}
	}
	newProduct := Product{ID: id, Name: name, Price: price, Stock: stock}
	inventory = append(inventory, newProduct)
	return nil
}

func updateStock(id int, newStock int) error {
	if newStock < 0 {
		return errors.New("stock cannot be negative")
	}
	for i, product := range inventory {
		if product.ID == id {
			inventory[i].Stock = newStock
			return nil
		}
	}
	return errors.New("product not found")
}

func searchProduct(identifier interface{}) (*Product, error) {
	for _, product := range inventory {
		switch v := identifier.(type) {
		case int:
			if product.ID == v {
				return &product, nil
			}
		case string:
			if product.Name == v {
				return &product, nil
			}
		}
	}
	return nil, errors.New("product not found")
}

func displayInventory() {
	fmt.Printf("%-5s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	fmt.Println("----------------------------------------------------")
	for _, product := range inventory {
		fmt.Printf("%-5d %-20s %-10.2f %-10d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

func sortInventory(choice string) {
	switch choice {
	case "price":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Price < inventory[j].Price
		})
	case "stock":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Stock < inventory[j].Stock
		})
	default:
		fmt.Println("Invalid sorting criteria. Use 'price' or 'stock'.")
	}
}

// Main function with a menu-driven program
func main() {
	inventory = []Product{
		{ID: 1, Name: "Laptop", Price: 50000, Stock: 2},
		{ID: 2, Name: "SmartPhone", Price: 40000, Stock: 5},
	}
	for {
		fmt.Println("\nInventory Management System")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Sort Inventory")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id, stock int
			var name string
			var price float64
			fmt.Print("Enter Product ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Product Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Product Price: ")
			fmt.Scan(&price)
			fmt.Print("Enter Stock: ")
			fmt.Scan(&stock)
			if err := addProduct(id, name, price, stock); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Product added successfully!")
			}

		case 2:
			var id, stock int
			fmt.Print("Enter Product ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter New Stock: ")
			fmt.Scan(&stock)
			if err := updateStock(id, stock); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Stock updated successfully!")
			}

		case 3:
			fmt.Println("Search by:")
			fmt.Println("1. ID")
			fmt.Println("2. Name")
			fmt.Print("Choose an option: ")
			var searchChoice int
			fmt.Scan(&searchChoice)

			if searchChoice == 1 {
				var id int
				fmt.Print("Enter Product ID: ")
				fmt.Scan(&id)
				product, err := searchProduct(id)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Printf("Found: %+v\n", *product)
				}
			} else if searchChoice == 2 {
				var name string
				fmt.Print("Enter Product Name: ")
				fmt.Scan(&name)
				product, err := searchProduct(name)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Printf("Found: %+v\n", *product)
				}
			} else {
				fmt.Println("Invalid choice.")
			}

		case 4:
			displayInventory()

		case 5:
			fmt.Println("Sort by:")
			fmt.Println("1. Price")
			fmt.Println("2. Stock")
			fmt.Print("Choose an option: ")
			var sortChoice int
			fmt.Scan(&sortChoice)

			if sortChoice == 1 {
				sortInventory("price")
			} else if sortChoice == 2 {
				sortInventory("stock")
			} else {
				fmt.Println("Invalid choice.")
			}
			displayInventory()

		case 6:
			fmt.Println("Exiting Inventory Management System. Goodbye!...")
			return

		default:
			fmt.Println("Invalid choice. Please choose a Valid Option from Menu.")
		}
	}
}
