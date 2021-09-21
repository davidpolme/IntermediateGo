package main

import "fmt"

type Employee struct {
	id   int
	name string
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
	e := Employee{}
	e.SetId(2)
	e.SetName("David")
	fmt.Printf("%v\n", e)
	fmt.Printf("id: %d, nombre: %s\n", e.GetId(), e.GetName())
}
