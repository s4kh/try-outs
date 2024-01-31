package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Suit int

const (
	spades Suit = iota + 1
	hearts
	diamonds
	clubs
)

func (s Suit) String() string {
	return [4]string{"spades", "hearts", "diamonds", "clubs"}[s-1]
}

type deck []card

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{spades.String(), hearts.String(), diamonds.String(), clubs.String()}
	cardValues := []string{"Ace", "King", "Queen", "Jack"}
	for i := 2; i <= 3; i++ {
		cardValues = append(cardValues, strconv.Itoa(i))
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, card{suit: suit, value: value})
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Since we need to modify the deck with remaining
// cards line:60 we pass the pointer
func (d *deck) deal(num int) deck {
	hand := (*d)[:num]
	*d = (*d)[num:]

	return hand
}

func (d deck) toString() string {
	s := make([]string, len(d))
	for _, val := range d {
		s = append(s, val.String())
	}
	return strings.Join(s, "\n")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newFromFile(name string) deck {
	// data, err := os.ReadFile(name)
	// if err != nil {
	// 	log.Fatalf("failed to read file: %v", err)
	// }

	// // str := string(data)
	// d := strings.Split(str, "\n")

	return newDeck()
}

func (d deck) shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

// this works because slice
// is passing the address
func (d deck) update() {
	fmt.Println("--------", &d)
	d[1] = card{suit: "ss", value: "ee"}
}
