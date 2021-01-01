package mw

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"
)

type Worker_Status struct {
	id     int
	status string
}

type Master_Data struct {
	Number_Workers    int
	Number_MapTask    int
	Number_ReduceTask int
	ws                []Worker_Status
}

type JOB int

func Start_Job(number_chuncks int) {

	//masterdata := Master_Data{"Some path", sync.WaitGroup{}}
	var wg sync.WaitGroup
	wg.Add(1)
	go StartServer(&wg)
	wg.Add(number_chuncks)
	go Start_Worker(number_chuncks, &wg)

	wg.Wait()

}

func (j *JOB) GetMapTask(message MessagePacket, reply *MessagePacket) error {

	masterDataStruct := Master_Data{0, 0, 0, []Worker_Status{}}
	masterDataStruct.Number_Workers += 1
	masterDataStruct.Number_MapTask += 1
	masterDataStruct.ws = append(masterDataStruct.ws, Worker_Status{message.Worker_id, "in-progress"})

	fmt.Println("GetMapTask called by", message.Worker_id)

	*reply = MessagePacket{message.Worker_id, "You will get a job soon..."}

	fmt.Println("Current list of workers and status")

	for _, val := range masterDataStruct.ws {
		fmt.Println(val.status, val.id)
	}

	//Send the path to the file chunks to each worker
	// Record how many maps are finished
	// Record workers as idle after finishing map task

	return nil
}

func StartServer(wg *sync.WaitGroup) {

	fmt.Println("Starting the master program..")

	master := new(JOB)

	err := rpc.Register(master)

	if err != nil {

		log.Fatal("Format of service Task isn't correct. ", err)

	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	fmt.Println("Serving RPC server on port 4040..", listener.Addr())

	error := http.Serve(listener, nil)

	fmt.Println("start working...")
	if error != nil {

		log.Fatal("Error serving: ", err)
	}

	fmt.Println("Waiting for client to request....")

	//wg.Done()

}
