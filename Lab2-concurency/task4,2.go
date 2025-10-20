package main

import (
	"fmt"
	"time"
)

func searchEngine(name string, delay time.Duration, query string, ch chan string) {
	time.Sleep(delay)
	ch <- fmt.Sprintf("Results from %s for '%s'", name, query)
}

func main() {
	fmt.Println("=== Multi Search Race ===")
	query := "go concurrency patterns"

	chA := make(chan string)
	chB := make(chan string)
	chC := make(chan string)

	go searchEngine("Engine A", 300*time.Millisecond, query, chA)
	go searchEngine("Engine B", 200*time.Millisecond, query, chB)
	go searchEngine("Engine C", 250*time.Millisecond, query, chC)

	start := time.Now()
	select {
	case result := <-chA:
		fmt.Printf("Engine A won! (%v)\n%s\n", time.Since(start), result)
	case result := <-chB:
		fmt.Printf("Engine B won! (%v)\n%s\n", time.Since(start), result)
	case result := <-chC:
		fmt.Printf("Engine C won! (%v)\n%s\n", time.Since(start), result)
	}
}
