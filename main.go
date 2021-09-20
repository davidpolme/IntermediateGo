package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)
	myvalue, err := strconv.ParseInt("NaN", 0, 64)
	if err != nil {
		fmt.Printf("Pero que ha pasao :o : %v\n", err)
	} else {
		fmt.Println("Todo bien, todo correcto. Y yo que me alegro :') :", myvalue)
	}

	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println(m["Key"])
	s := []string{"Abra", "Ka", "Dabra"}
	for index, value := range s {
		fmt.Printf("%d : %s\n", index, value)
	}
	s = append(s, ":)")
	fmt.Println("___________________________________")
	for index, value := range s {
		fmt.Printf("%d : %s\n", index, value)
	}

	c := make(chan int)
	go doSomething(c)
	<-c

	g := 25

	fmt.Println(g)

	h := &g

	fmt.Println(*h)
}

func doSomething(c chan int) {
	time.Sleep(time.Second * 1)
	fmt.Println("Done")
	c <- 1
}
