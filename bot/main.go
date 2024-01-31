package main

import "fmt"

func printGreeting(greetingBot greeting) {
	fmt.Println(greetingBot.getGreeting())
}

func main() {
	fmt.Println("---bot speaking---")
	englishBot := bot{lang: "en"}
	spanishBot := bot{lang: "spain"}

	printGreeting(englishBot)
	printGreeting(spanishBot)
}
