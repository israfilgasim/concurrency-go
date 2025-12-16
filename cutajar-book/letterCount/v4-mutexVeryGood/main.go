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
	time.Sleep(10 * time.Second)
	mutex.Lock()
	for i, c := range AllLetters {
		fmt.Printf("%c-%d ", c, frequency[i])
	}
	mutex.Unlock()
}

func CountLetters(url string, frequency []int, mutex *sync.Mutex) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", url, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server returning error status code: " + resp.Status)
	}
	current := make(map[int]int, 26)
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(AllLetters, c)
		if cIndex >= 0 {
			current[cIndex] += 1
		}
	}
	mutex.Lock()
	for k, v := range current {
		frequency[k] += v
	}
	mutex.Unlock()
	fmt.Println("Completed:", url, time.Now().Format("15:04:05"))
}

/*
By collecting the letter counts in a local map first and then updating the shared frequency slice while holding the mutex lock, we minimize the time spent holding the lock. This allows other goroutines to proceed with their work concurrently, improving overall performance.
*/
