package main

import (
	"fmt"
)

func goroutine(quit chan bool) {
	for {
		select {
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("Do your thing")
		}
	}
}

// Go routine will still be running,
// after you return from this function.
func invoker() {
	q := make(chan bool)
	go goroutine(q)

	q <- true
}
func main() {

}
