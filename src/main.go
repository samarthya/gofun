package main

import (
	"fmt"
	"time"
)

func send(ch chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			ch <- i
		}

		// os.Exit(0)
		fmt.Println(" Sender done...")
		// Close the channel
		close(ch)
	}()
}

func receive(ch chan int) {

	go func() {
		defer func() { done <- true }()

		for val := range ch {

			if val >= 9 {
				fmt.Println(" Closing reciever...")
				fmt.Printf(" > %d\n", val)
				// os.Exit(0)
				break
			}

			fmt.Printf(" > %d\n", val)
		}
		fmt.Println(" Receiver done...")
	}()

}

var done chan bool = make(chan bool)

func main() {
	fmt.Println(" --- Main ---")
	ch := make(chan int)

	send(ch)
	receive(ch)

	// wait on done
	<-done
	fmt.Println(" Done.. ")
}
