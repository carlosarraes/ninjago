package main

import "fmt"

func main() {
	x := 2 == 3
	y := 2 <= 3
	z := 4 >= 2
	a := 2 != 5
	b := 3 < 2
	c := 2 > 4

	fmt.Println(x, y, z, a, b, c)
}
