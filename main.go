package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTymeEmployee struct {
	Employee
	Person
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	//SELECT * FROM Persona WHERE ...
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	//
	return Employee{}, nil
}

var GetFullTimeEmployeeByID = func(id int, dni string) (FullTymeEmployee, error) {
	var ftEmployee FullTymeEmployee
	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Employee = e

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Person = p

	return ftEmployee, nil
}

func main() {
}
