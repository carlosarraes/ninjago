package main

import "fmt"

type hot int

func main() {
	var x hot
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
