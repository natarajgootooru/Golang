package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func (e *employee) printDetails() {

	fmt.Printf("Name: %s | Age: %d | Salary: %d\n", e.name, e.age, e.salary)
}

func (e *employee) updateSalary(updatedSalary int) {

	e.salary = updatedSalary
}
