package main

import "testing"

func TestSum(t *testing.T) {
	/*
		total := Sum(5, 5)
		expectedTotal := 10
		if total != expectedTotal {
			t.Errorf("Sum was incorrect, got %d expected %d", total, expectedTotal)
		}
	*/
	//Tabla con los test cases

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrecct, got %d expected %d", total, item.n)
		}
	}
}
