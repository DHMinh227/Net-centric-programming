package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Student struct {
	ID         int
	StudyHours int
}

func student(id int, studyHours int, library chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	library <- true // try enter library
	fmt.Printf("Student %d entered library, will study for %d hours\n", id, studyHours)

	time.Sleep(time.Duration(studyHours) * time.Second)

	<-library // leave
	fmt.Printf("Student %d left library after %d hours\n", id, studyHours)
}

func main() {
	fmt.Println("=== Library Simulation ===")
	fmt.Println("Library capacity: 30 students")
	fmt.Println("Total students today: 100")
	fmt.Println("Simulation: 1 second = 1 hour\n")

	rand.Seed(time.Now().UnixNano())
	library := make(chan bool, 30)
	var wg sync.WaitGroup
	start := time.Now()

	students := make([]Student, 100)
	for i := 0; i < 100; i++ {
		students[i] = Student{
			ID:         i + 1,
			StudyHours: rand.Intn(4) + 1,
		}
	}

	for _, s := range students {
		wg.Add(1)
		go student(s.ID, s.StudyHours, library, &wg)
	}

	wg.Wait()
	duration := time.Since(start)

	fmt.Println("\n=== Simulation Complete ===")
	fmt.Printf("Total students served: %d\n", len(students))
	fmt.Printf("Library was open for: %.1f hours\n", duration.Seconds())
}
