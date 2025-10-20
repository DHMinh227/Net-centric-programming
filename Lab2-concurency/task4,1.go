package main

import (
	"fmt"
	"time"
)

func searchEngineA(query string, ch chan string) {
	time.Sleep(300 * time.Millisecond)
	ch <- fmt.Sprintf("Results from Engine A for '%s'", query)
}

func searchEngineB(query string, ch chan string) {
	time.Sleep(200 * time.Millisecond)
	ch <- fmt.Sprintf("Results from Engine B for '%s'", query)
}

func main() {
	fmt.Println("=== Search Race ===")
	query := "golang concurrency"

	chA := make(chan string)
	chB := make(chan string)

	go searchEngineA(query, chA)
	go searchEngineB(query, chB)

	start := time.Now()
	select {
	case result := <-chA:
		fmt.Printf("Engine A won! (%v)\n%s\n", time.Since(start), result)
	case result := <-chB:
		fmt.Printf("Engine B won! (%v)\n%s\n", time.Since(start), result)
	}
}
