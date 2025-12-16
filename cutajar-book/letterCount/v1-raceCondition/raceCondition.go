package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int) {
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
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	fmt.Println("Completed:", url)
}

func main() {
	var sequentialFrequency = make([]int, 26)
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		countLetters(url, sequentialFrequency)
	}
	for i, c := range allLetters {
		fmt.Printf("%c-%d ", c, sequentialFrequency[i])
	}

	var parallelFrequency = make([]int, 26)
	for i := 1010; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, parallelFrequency)
	}
	// Wait for a while to let goroutines finish
	time.Sleep(10 * time.Second)
	for i, c := range allLetters {
		fmt.Printf("%c-%d ", c, parallelFrequency[i])
	}
}
