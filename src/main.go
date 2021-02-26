package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	os.Exit(0)
}

func receive(ch chan int) {
	var val int

	for {
		val = <-ch
		fmt.Printf(" > %d\n", val)
		if val >= 10 {
			close(ch)
			done <- true
			os.Exit(0)
		}
	}
}

var done chan bool = make(chan bool)

func main() {
	fmt.Println(" --- Main ---")
	env := os.Environ()
	goroot := runtime.GOROOT()
	//Use sysctl -n hw.ncpu for cores
	fmt.Printf(" GOROOT: %s\n", goroot)

	runtime.GOMAXPROCS(16)
	for i, v := range env {
		if strings.Contains(v, "GOMAXPROCS") {
			fmt.Printf(" Found %v at %d", v, i)
		}
	}

	ch := make(chan int)

	go send(ch)
	go receive(ch)

	// wait on done
	<-done
	fmt.Println(" Done.. ")
}
