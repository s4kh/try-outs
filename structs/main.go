package main

import "fmt"

type celcius float64

func (c celcius) String() string {
	return fmt.Sprintf("%.2fC", c)
}

type temperature struct {
	celcius
}

func (t temperature) String() string {
	return fmt.Sprintf("Current temp %.3fC", t)
}

func main() {
	c := celcius(14.0)
	fmt.Println(c)

	t := temperature{c}
	fmt.Println(t)
}
