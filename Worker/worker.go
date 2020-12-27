package worker

import "fmt"

type KeyValue struct {
	Key   string
	Value string
}

func Hello() {

	fmt.Println("From the worker")
}
