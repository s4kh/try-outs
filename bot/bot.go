package main

import "fmt"

type greeting interface {
	getGreeting() string
}

type bot struct {
	lang string
}

func (b bot) getGreeting() string {
	return fmt.Sprintf("%s greeting", b.lang)
}
