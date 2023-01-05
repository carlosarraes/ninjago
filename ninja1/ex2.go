package main

import "fmt"

var (
	x int
	y string
	z bool
)

func main() {
	fmt.Println(x, y, z)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
}
