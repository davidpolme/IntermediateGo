package main

import "fmt"

//canal exclusivo de escritura <- a la derecha de chan
func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

//in: canal exclusivo de lectura <- a la izquierda de chan
//out: canal exclusivo de escritura <- a la derecha de chan
func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- 2 * value
		//in <- 1
	}
	close(out)
}

func PrinntValues(c chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Double(generator, doubles)
	PrinntValues(doubles)

}
