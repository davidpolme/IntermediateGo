package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		fmt.Println("Starting Function")
		time.Sleep(2 * time.Second)
		fmt.Println("End")
		c <- 1
	}()
	<-c
}
