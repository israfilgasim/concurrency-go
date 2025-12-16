package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func stingy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*money += 10
		mutex.Unlock()
	}
	fmt.Println("Stingy Done")
}

func spendy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 200000; i++ {
		mutex.Lock()
		for *money < 50 {
			mutex.Unlock()
			time.Sleep(10 * time.Millisecond)
			mutex.Lock()
		}
		*money -= 50
		if *money < 0 {
			fmt.Println("Money is negative!")
			os.Exit(1)
		}
		mutex.Unlock()
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100
	mutex := sync.Mutex{}
	go stingy(&money, &mutex)
	go spendy(&money, &mutex)
	time.Sleep(2 * time.Second)
	mutex.Lock()
	fmt.Println("Money in bank account: ", money)
	mutex.Unlock()

}

/*
This solution will work for our use case, but it’s not ideal. In our example, we choose the arbitrary sleep value of 10 milliseconds, but what would be the optimal number to choose? At one extreme, we can choose not to sleep at all. This ends up wasting CPU resources, as the CPU would be cycling needlessly, checking the money variable even if it doesn’t change. At the other extreme, if the goroutine sleeps for too long, we might waste time waiting for a change in the money variable that has already happened.
*/
