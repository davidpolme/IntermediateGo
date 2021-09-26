package main

import "testing"

func TestGetFulltimeEmployeeByID(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTymeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						Name: "John Doe",
						Age:  34,
						DNI:  "1",
					}, nil
				}
			},
		},
	}
}
