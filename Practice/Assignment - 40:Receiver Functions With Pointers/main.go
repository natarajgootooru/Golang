package main

func main() {

	e1 := employee{name: "Gautham", age: 29, salary: 20000}
	e1.printDetails()
	e1.updateSalary(30000)
	e1.printDetails()
}
