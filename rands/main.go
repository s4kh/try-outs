package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Println(x, y)

	switch {
	case x < 4 && y < 4:
		fmt.Println("both are less than 4")
	case x > 6 && y > 6:
		fmt.Println("both are greater than 6")
	default:
		fmt.Println("not handled!")
	}
}
