package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

const AllLetters = "abcdefghijklmnopqrstuvwxyz"

func main() {
	mutex := sync.Mutex{}
	var frequency = make([]int, 26)
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go CountLetters(url, frequency, &mutex)
	}
	time.Sleep(60 * time.Second)
	mutex.Lock()
	for i, c := range AllLetters {
		fmt.Printf("%c-%d ", c, frequency[i])
	}
	mutex.Unlock()
}

// CountLetters
// Note: this program us locking the entire goroutine with mutex on purpose to demonstrate
// bad placement of the lock and unlock. We fix this in the next listing
func CountLetters(url string, frequency []int, mutex *sync.Mutex) {
	mutex.Lock()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", url, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server returning error status code: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(AllLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	fmt.Println("Completed:", url, time.Now().Format("15:04:05"))
	mutex.Unlock()
}

/*
By using mutexes in this manner, we have changed our concurrent program into a sequential one. We will end up downloading and processing one web page at a time since we’re needlessly blocking the entire execution. If we go ahead and run this, the time taken will be the same as the non-concurrent version of the program, although the order of execution will be random
*/
