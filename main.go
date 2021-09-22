package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporatyEmployee struct {
	Person
	Employee
	taxRate int
}

func (tempEmployee TemporatyEmployee) getMessage() string {
	return "Temporary Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "David"
	ftEmployee.age = 24
	ftEmployee.id = 1
	getMessage(ftEmployee)

	tempEmployee := TemporatyEmployee{}
	tempEmployee.name = "Tatiana"
	tempEmployee.age = 25
	tempEmployee.id = 2
	getMessage(tempEmployee)
}
