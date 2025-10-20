package main

import (
	"fmt"
	"sync"
	"time"
)

func findEvens(numbers []int, result chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var evens []int
	for _, n := range numbers {
		if n%2 == 0 {
			evens = append(evens, n)
		}
		time.Sleep(50 * time.Millisecond)
	}
	result <- evens
}

func findOdds(numbers []int, result chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var odds []int
	for _, n := range numbers {
		if n%2 != 0 {
			odds = append(odds, n)
		}
		time.Sleep(50 * time.Millisecond)
	}
	result <- odds
}

func findSquares(numbers []int, result chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var squares []int
	for _, n := range numbers {
		squares = append(squares, n*n)
		time.Sleep(50 * time.Millisecond)
	}
	result <- squares
}

func main() {
	fmt.Println("=== Number Processor ===")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenChan := make(chan []int)
	oddChan := make(chan []int)
	squareChan := make(chan []int)

	var wg sync.WaitGroup
	wg.Add(3)

	start := time.Now()

	go findEvens(numbers, evenChan, &wg)
	go findOdds(numbers, oddChan, &wg)
	go findSquares(numbers, squareChan, &wg)

	go func() {
		wg.Wait()
		close(evenChan)
		close(oddChan)
		close(squareChan)
	}()

	evens := <-evenChan
	odds := <-oddChan
	squares := <-squareChan

	fmt.Println("Numbers:", numbers)
	fmt.Println("Evens:", evens)
	fmt.Println("Odds:", odds)
	fmt.Println("Squares:", squares)
	fmt.Printf("\nTime: %s\n", time.Since(start))
}
