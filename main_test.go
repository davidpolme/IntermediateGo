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

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{5, 1, 5},
		{2, 5, 5},
		{1, 10, 10},
	}
	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("GetMax was incorrect, got %d, expected %d", max, item.n)
		}
	}
}

func TestFib(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}
	for _, item := range tables {
		fib := Fibonnaci(item.a)
		if fib != item.n {
			t.Errorf("Fibonacci was incorrect, got %d expected %d", fib, item.n)
		}
	}
}
