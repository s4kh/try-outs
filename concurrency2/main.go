package main

import (
	"fmt"
	"sync"
	"time"
)

func attack(e string) {
	fmt.Println("Attacked ", e)
}

func main() {
	var wg sync.WaitGroup

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	enemies := []string{"Kai", "Rui", "Moody", "Rauto", "Boonjai"}
	wg.Add(len(enemies))

	for _, enemy := range enemies {
		e := enemy
		go func() {
			attack(e)
			wg.Done()
		}()
	}

	wg.Wait()

}
