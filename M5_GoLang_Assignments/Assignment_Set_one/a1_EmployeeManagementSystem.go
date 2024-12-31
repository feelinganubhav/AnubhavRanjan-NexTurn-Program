/*
Exercise 1: Employee Management System
Topics Covered: Go Conditions, Go Loops, Go Constants, Go Functions, Go Arrays, Go
Strings, Go Errors
Case Study:
A company wants to manage its employees' data in memory. Each employee has an ID,
name, age, and department. You need to build a small application that performs the
following:
1. Add Employee: Accept input for employee details and store them in an array of
structs. Validate the input:
o ID must be unique.
o Age should be greater than 18. If validation fails, return custom error
messages.
2. Search Employee: Search for an employee by ID or name using conditions.
Return the details if found, or return an error if not found.
3. List Employees by Department: Use loops to filter and display all employees in
a given department.
4. Count Employees: Use constants to define a department (e.g., "HR", "IT"), and
display the count of employees in that department.
Bonus:
Refactor the repetitive code using functions, and add error handling for invalid
operations like searching for a non-existent employee.
*/

package main

import (
	"errors"
	"fmt"
	"strings"
)

type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

const (
	HR  = "HR"
	IT  = "IT"
	FIN = "Finance"
)

var employees []Employee

func addEmployee(id int, name string, age int, department string) error {
	if age <= 18 {
		return errors.New("age must be greater than 18")
	}

	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("ID must be unique")
		}
	}

	employees = append(employees, Employee{ID: id, Name: name, Age: age, Department: department})
	return nil
}

func searchEmployee(query string) (*Employee, error) {
	for _, emp := range employees {
		if fmt.Sprintf("%d", emp.ID) == query || strings.EqualFold(emp.Name, query) {
			return &emp, nil
		}
	}
	return nil, errors.New("employee not found")
}

func listEmployeesByDepartment(department string) []Employee {
	var result []Employee
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			result = append(result, emp)
		}
	}
	return result
}

func countEmployees(department string) int {
	count := 0
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("Employee Management System")

	// Adding employees
	err := addEmployee(1, "Alice", 30, HR)
	if err != nil {
		fmt.Println("Error adding employee:", err)
	}

	err = addEmployee(2, "Bob", 25, IT)
	if err != nil {
		fmt.Println("Error adding employee:", err)
	}

	err = addEmployee(3, "Charlie", 28, HR)
	if err != nil {
		fmt.Println("Error adding employee:", err)
	}

	fmt.Println("Some Employees added successfully!")

	for {
		fmt.Println("\nEmployee Management System")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. List Employees by Department")
		fmt.Println("4. Count Employees in a Department")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 5.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Adding a new employee...")
			fmt.Print("Enter ID: ")
			var id int
			fmt.Scan(&id)

			fmt.Print("Enter Name: ")
			var name string
			fmt.Scan(&name)

			fmt.Print("Enter Age: ")
			var age int
			fmt.Scan(&age)

			fmt.Print("Enter Department: ")
			var department string
			fmt.Scan(&department)

			err := addEmployee(id, name, age, department)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Employee added successfully.")
			}
		case 2:
			fmt.Println("Search Employee")
			fmt.Print("Please Enter a Valid ID or Name: ")
			var searchChoice string
			fmt.Scan(&searchChoice)

			employee, err := searchEmployee(searchChoice)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Employee Found: %+v\n", employee)
			}

		case 3:
			fmt.Println("List Employees by Department")
			fmt.Print("Enter Department (E.g. HR, IT, Finance): ")
			var department string
			fmt.Scan(&department)
			employs := listEmployeesByDepartment(department)

			if len(employs) == 0 {
				fmt.Printf("Error: No Employees Found in %s Department\n:", department)
			} else {
				// Listing employees by department
				fmt.Printf("Employees in %s Department\n:", department)
				for _, emp := range employs {
					fmt.Printf("%+v\n", emp)
				}
			}

		case 4:
			fmt.Println("Count Employees in a Department")
			fmt.Print("Enter Department (E.g. HR, IT, Finance): ")
			var department string
			fmt.Scan(&department)
			count := countEmployees(department)

			fmt.Printf("Number of employees in %s: %d\n", department, count)

		case 5:
			fmt.Println("Exiting Employee Management System. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
