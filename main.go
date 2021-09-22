package main

import (
	"fmt"

	"github.com/donvito/hellomod"

	himodV2 "github.com/donvito/hellomod/v2"
)

func main() {
	fmt.Println("Hello World!")
	hellomod.SayHello()
	himodV2.SayHello("Holaas")
}
