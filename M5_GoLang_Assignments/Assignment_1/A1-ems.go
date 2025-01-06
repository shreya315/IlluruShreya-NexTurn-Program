package main

import (
	"errors"
	"fmt"
)

// Employee struct to store employee details
type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

// Constants for department names
const (
	HR      = "HR"
	IT      = "IT"
	Finance = "Finance"
)

// Slice to store employees
var employees []Employee

// Function to add an employee
func addEmployee(id int, name string, age int, department string) error {
	// Check if ID is unique
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("employee ID must be unique")
		}
	}

	// Validate age
	if age <= 18 {
		return errors.New("employee age must be greater than 18")
	}

	// Add employee to the list
	employees = append(employees, Employee{ID: id, Name: name, Age: age, Department: department})
	return nil
}

// Function to search for an employee by ID or name
func searchEmployee(id int, name string) (*Employee, error) {
	for _, emp := range employees {
		if emp.ID == id || emp.Name == name {
			return &emp, nil
		}
	}
	return nil, errors.New("employee not found")
}

// Function to list employees by department
func listEmployeesByDepartment(department string) []Employee {
	var result []Employee
	for _, emp := range employees {
		if emp.Department == department {
			result = append(result, emp)
		}
	}
	return result
}

// Function to count employees in a department
func countEmployees(department string) int {
	count := 0
	for _, emp := range employees {
		if emp.Department == department {
			count++
		}
	}
	return count
}

func main() {
	// Adding employees
	err := addEmployee(1, "Alice", 25, IT)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = addEmployee(2, "Bob", 30, HR)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = addEmployee(3, "Charlie", 22, IT)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Searching for an employee
	emp, err := searchEmployee(2, "")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Employee Found:", *emp)
	}

	// Listing employees by department
	fmt.Println("Employees in IT:")
	for _, emp := range listEmployeesByDepartment(IT) {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Department: %s\n", emp.ID, emp.Name, emp.Age, emp.Department)
	}

	// Counting employees in HR
	fmt.Printf("Number of employees in HR: %d\n", countEmployees(HR))
}
