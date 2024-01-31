package main

import "fmt"

func print(pointerToString *string) {
	fmt.Println(&pointerToString)
}

func main() {
	name := "Bill"

	namePointer := &name
	fmt.Println(namePointer)
	print(namePointer)

	// c := card{suit: clubs.String(), value: "Ace"}

	// Turn address(pointer variable) in to a value by
	// *c
	// Turn value into an address by
	// &c
	// fmt.Println(c, &c)

	// c.updateSuit(hearts.String())
	// fmt.Println(c)

	// d := newDeck()
	// d.print()
	// d.update()
	// fmt.Println(d[2])
	// d.print()
	// hand := d.deal(2)
	// fmt.Println("---Hand---")
	// hand.print()
	// fmt.Println("---Deck after deal---")
	// d.print()
	// fmt.Println("--------", d)
	// hand.update()
	// fmt.Println("---After update---")
	// hand.print()
	// fmt.Println("---After shuffle---")
	// hand.shuffle()
	// hand.print()

	// fmt.Println("---start of the hand deck---")
	// hand.print()
	// fmt.Println("---hand finished---")
	// // d.print()

	// if err := d.saveToFile("my_deck"); err != nil {
	// 	log.Fatalf("failed to save to file: %v", err)
	// }
	// hand.saveToFile()
	// d2 := newFromFile("my_deck2")

	// d2.print()

}
