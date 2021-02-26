package main

import (
	"fmt"
	"time"
)

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		ch <- i
	}

	// os.Exit(0)
	fmt.Println(" Sender done...")
	return
}

func term() {
	done <- true
}

func receive(ch chan int) {
	var val int

	defer term()

	for {

		if val = <-ch; val >= 9 {
			fmt.Println(" Closing reciever...")
			fmt.Printf(" > %d\n", val)

			// Close the channel
			close(ch)
			// os.Exit(0)
			break
		}

		fmt.Printf(" > %d\n", val)
	}

	fmt.Println(" Receiver done...")
}

var done chan bool = make(chan bool)

func main() {
	fmt.Println(" --- Main ---")
	ch := make(chan int)

	go send(ch)
	go receive(ch)

	// wait on done
	<-done
	fmt.Println(" Done.. ")
}
