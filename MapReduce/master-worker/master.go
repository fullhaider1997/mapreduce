package mw

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
	"strings"
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

var master = Master_Data{
	Number_Workers:    0,
	Number_MapTask:    0,
	Number_ReduceTask: 0,
	ws:                []Worker_Status{},
}

func Start_Job(number_chuncks int) {

	var wg sync.WaitGroup
	//Initalize master data structure based on number of chunks
	master = master.init(number_chuncks)
	fmt.Println("The number of workers", master.Number_Workers)
	wg.Add(1)
	go StartServer(&wg)
	wg.Add(number_chuncks)
	go Start_Worker(number_chuncks, &wg)

	wg.Wait()

}

var path string

func getPath() string {

	return path
}

func setPath(pathd string) {
	path = pathd
}

func (j *JOB) GetMapTask(message MessagePacket, reply *MessagePacket) error {

	fmt.Println("GetMapTask called by", message.Worker_id)

	fmt.Println("The number of workers", master.Number_Workers)

	for i := 0; i < master.Number_Workers; i++ {

		if strings.Compare(master.ws[i].status, "idle") == 0 {

			fmt.Println("Worker is idle !")

			*reply = MessagePacket{message.Worker_id, getPath() + "/chunk-" + strconv.Itoa(i) + ".txt"}

		} else {

			*reply = MessagePacket{message.Worker_id, "We didn't get the path"}

		}

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

////////////// Master data //////////////
func (m Master_Data) init(num int) Master_Data {

	m.Number_Workers = num

	for i := 0; i < num; i++ {
		m.ws = append(m.ws, Worker_Status{0, "idle"})
	}
	return m
}

func (m Master_Data) addWorker(id int) {

	m.Number_Workers += 1
	m.Number_MapTask += 1

	m.ws = append(m.ws, Worker_Status{id, "idle"})

}

func (m Master_Data) getMasterDSWorkers() []Worker_Status {

	return m.ws
}

func (m Master_Data) getMasterDSNumWorkers() int {

	return m.Number_Workers
}
