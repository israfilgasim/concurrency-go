package main

import (
	"fmt"
	"time"
)

func sequentialDoWork(id int) {
	fmt.Printf("Sequential Work %d started at %s\n", id, time.Now().Format("15:04:05"))
	time.Sleep(1 * time.Second)
	fmt.Printf("Sequiential Work %d finished at %s\n", id, time.Now().Format("15:04:05"))
}

func parallelDoWork(id int) {
	fmt.Printf("Parallel Work %d started at %s\n", id, time.Now().Format("15:04:05"))
	time.Sleep(1 * time.Second)
	fmt.Printf("Parallel Work %d finished at %s\n", id, time.Now().Format("15:04:05"))
}

func main() {
	fmt.Println("Starting sequential work...")
	for i := 0; i < 5; i++ {
		sequentialDoWork(i)
	}
	fmt.Println("----------------------------")

	fmt.Println("Starting parallel work...")
	for i := 0; i < 5; i++ {
		go parallelDoWork(i)
	}
	// Wait for a while to let goroutines finish
	time.Sleep(2 * time.Second)

}

// Output:
/*
Starting sequential work...
Sequential Work 0 started at 20:53:53
Sequiential Work 0 finished at 20:53:54
Sequential Work 1 started at 20:53:54
Sequiential Work 1 finished at 20:53:55
Sequential Work 2 started at 20:53:55
Sequiential Work 2 finished at 20:53:56
Sequential Work 3 started at 20:53:56
Sequiential Work 3 finished at 20:53:57
Sequential Work 4 started at 20:53:57
Sequiential Work 4 finished at 20:53:58
----------------------------
Starting parallel work...
Parallel Work 4 started at 20:53:58
Parallel Work 0 started at 20:53:58
Parallel Work 2 started at 20:53:58
Parallel Work 1 started at 20:53:58
Parallel Work 3 started at 20:53:58
Parallel Work 3 finished at 20:53:59
Parallel Work 1 finished at 20:53:59
Parallel Work 4 finished at 20:53:59
Parallel Work 0 finished at 20:53:59
Parallel Work 2 finished at 20:53:59
*/
