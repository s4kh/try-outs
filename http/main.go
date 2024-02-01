package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type customWriter struct{}

func (customWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	fmt.Printf("Just wrote %d bytes to the terminal\n", len(p))

	return len(p), nil
}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		log.Fatalln("error during get:", err)
	}

	cw := customWriter{}

	io.Copy(cw, resp.Body)
}
