package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d Started \n", id)
	fmt.Printf("Worker %d End \n", id)
	wg.Done()
}

func main() {
	fmt.Println("Learning Sync Wait Group...")

	var wg sync.WaitGroup
	// Start 3 goroutines
	for i:=1; i<=3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for all 3 goroutines to finish
	wg.Wait()
	fmt.Println("All workers completed...")
}
