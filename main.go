package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

//4
func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	//1
	e := Employee{}
	fmt.Printf("%v\n", e)

	//2
	e2 := Employee{
		id:       1,
		name:     "Tatiana",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)
	//3
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 1
	e3.name = "Odi"
	e3.vacation = false
	fmt.Printf("%v\n", *e3)

	//4
	e4 := NewEmployee(1, "Pepe", false)
	fmt.Printf("%v\n", *e4)
}
