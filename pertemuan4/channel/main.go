package main

import (
	"fmt"
	"sync"
)

// Channel
// merukapan sebuah media untuk gorotine dapat berkomunikasi/bertukar data
func main() {

	wg := sync.WaitGroup{}

	channel := make(chan int)

	wg.Add(2)
	go receiveData(channel, 1, &wg)
	go receiveData(channel, 2, &wg)

	for i := 0; i < 10; i++ {
		// mengirim data ke sebuah channel
		channel <- i
	}

	close(channel)
	wg.Wait()
}

// fungsi untuk mengirim data
func receiveData(channel chan int, number int, wg *sync.WaitGroup) {
	for data := range channel {
		fmt.Println("Data received ", data, "Form Gorotine Number", number)
	}
	wg.Done()
}
