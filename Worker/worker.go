package worker

import (
	master "MapReduce/Master"
	"fmt"
	"log"
	"net/rpc"
)

type KeyValue struct {
	Key   string
	Value string
}

type Worker_Data struct {
	worker_id int
}

func Start_Worker(numOfWorkers int) {

	for i := 0; i < numOfWorkers; i++ {

		wd := Worker_Data{i}
		go Task(&wd)

	}
}

type fileStruct master.FileData

func Task(wd *Worker_Data) {

	filestruct := fileStruct{"dsdsdsds"}
	var reply fileStruct

	fmt.Println("Worker id is ", *&wd.worker_id)

	// Create a TCP connection to localhost on port 1234
	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	fmt.Println("Client is connect through port ")

	client.Call("Job.GetTaskFromMaster", filestruct, &reply)

}
