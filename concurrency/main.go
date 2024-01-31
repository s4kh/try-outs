package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	t1 := rand.Intn(255)
	t2 := rand.Intn(255)

	ch1 := make(chan int)
	ch2 := make(chan int)
	// wg.Add(1)
	go func() {
		// defer wg.Done()
		d1 := time.Duration(t1)
		time.Sleep(d1 * time.Millisecond)
		ch1 <- t1
	}()
	// wg.Add(1)
	go func() {
		// defer wg.Done()
		d2 := time.Duration(t2)
		time.Sleep(d2 * time.Millisecond)
		ch2 <- t2
	}()

	for i := 0; i < 2; i++ {
		select {
		case g1 := <-ch1:
			fmt.Println("value from routine 1", g1)
		case g2 := <-ch2:
			fmt.Println("value from routine 2", g2)
		}
	}

	// wg.Wait()

}
