package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", i)
	//Some task is happening
	fmt.Printf("Worker %d end\n", i)
}

//worker(1)
//worker(2)
//worker(3)

func main() {
	//fmt.Println("Exploring Wait Group")
	var wg sync.WaitGroup //tasklist
	//start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) //Increment the WaitGroup counter
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("Workers task Completed")
}


// Output :-

// Worker 1 started
// Worker 1 end
// Worker 3 started
// Worker 3 end
// Worker 2 started
// Worker 2 end
// Workers task Completed