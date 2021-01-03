package mw

import (
	wordcount "User-mapreduce/MapReduce/wordcount"
	"fmt"
	"io/ioutil"
	"log"
	"net/rpc"
	"sync"
)

type MessagePacket struct {
	Worker_id int
	Message   string
}

func Start_Worker(numOfWorkers int, wg *sync.WaitGroup) {

	for i := 0; i < numOfWorkers; i++ {

		wp := MessagePacket{i, "Waiting for a job master"}
		go Task(&wp, wg)

	}
}

var index int = 0

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

	chunkPath := reply.Message

	content, err := ioutil.ReadFile(chunkPath)

	if err != nil {
		log.Fatal(err)
	}

	kv := wordcount.Mapf(chunkPath, string(content))

	if kv != nil {

		wp.Message = "complete"

		client.Call("JOB.NotifyMaster", wp, &reply)

	} else {

		wp.Message = "fail"

		client.Call("JOB.NotifyMaster", wp, &reply)
	}

}
