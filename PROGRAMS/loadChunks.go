package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

const mb = 1024 * 1024
const gb = 1024 * mb

func main() {
	// A waitgroup to wait for all go-routines to finish.
	wg := sync.WaitGroup{}

	// This channel is used to send every read word in various go-routines.
	channel := make(chan (string))

	// A dictionary which stores the count of unique words.
	dict := make(map[string]int64)

	// Done is a channel to signal the main thread that all the words have been
	// entered in the dictionary.
	done := make(chan (bool), 1)

	// Read all incoming words from the channel and add them to the dictionary.
	go func() {
		for s := range channel {
			fmt.Println(s)
			dict[s]++
		}

		// Signal the main thread that all the words have entered the dictionary.
		done <- true
	}()

	// Current signifies the counter for bytes of the file.
	var current int64

	// Limit signifies the chunk size of file to be processed by every thread.
	var limit int64 = 15 * mb

	for i := 0; i < 1; i++ {
		wg.Add(1)

		go func() {
			read(current, limit, "gameofthrones.txt", channel)
			fmt.Printf("%d thread has been completed \n", i)
			wg.Done()
		}()

		// Increment the current by 1+(last byte read by previous thread).
		current += limit + 1
	}

	// Wait for all go routines to complete.
	wg.Wait()
	close(channel)

	// Wait for dictionary to process all the words.
	<-done
	close(done)
	fmt.Println(dict)
}

func read(offset int64, limit int64, fileName string, channel chan (string)) {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Move the pointer of the file to the start of designated chunk.
	// file.Seek(offset, 0)
	r := bufio.NewScanner(file)
	r.Split(bufio.ScanWords)
	for r.Scan() {
		s := r.Text()
		if s != "" {
			// Send the read word in the channel to enter into dictionary.
			channel <- s
		}
	}
}
