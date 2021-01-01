package mw

import (
	"fmt"
	"log"
	"net/rpc"
	"sync"
)

type KeyValue struct {
	Key   string
	Value string
}

type MessagePacket struct {
	Worker_id int
	Message   string
}

func Start_Worker(numOfWorkers int, wg *sync.WaitGroup) {

	for i := 0; i < numOfWorkers; i++ {

		wd := MessagePacket{i, "Waiting for a job master"}
		go Task(&wd, wg)

	}
}

func Task(wp *MessagePacket, wg *sync.WaitGroup) {

	fmt.Println("Running workerWorker id is ", *&wp.Worker_id)

	var reply MessagePacket
	// Create a TCP connection to localhost on port 1234
	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	fmt.Println("Client is connect through port... ")

	client.Call("JOB.GetMapTask", wp, &reply)

	fmt.Println("Message from the server..")
	fmt.Println(reply.Message)

}
