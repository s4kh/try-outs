package main

import "fmt"

type card struct {
	suit  string
	value string
}

func (c card) String() string {
	return fmt.Sprintf("%s of %s", c.value, c.suit)
}

func (c *card) updateSuit(suit string) {
	c.suit = suit
}
