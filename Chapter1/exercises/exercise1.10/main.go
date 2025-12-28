// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	filename := "result.txt"

	file, err := os.OpenFile(filename, os.O_CREATE, 0644)

	if err != nil {
		panic("An error occured while openning file!")
	}

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // Start a goroutine
	}

	for range os.Args[1:] {
		//fmt.Println(<-ch) // Receive from a channel ch
		fmt.Fprintln(file, <-ch)
	}

	totalTime := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Fprintln(file, totalTime)

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // Send to channel ch
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // Don't leak resources

	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
