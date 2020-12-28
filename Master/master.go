package master

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type FileData struct {
	FileNamePath string
}

type Job int

func Start() {

	go StartServer()

}

func (J *Job) GetTaskFromMaster(filePath FileData, reply *FileData) error {

	fmt.Println("Calling the server to ask for a task")

	*reply = filePath

	return nil
}

func StartServer() {

	fmt.Println("Starting the master program..")

	master := new(Job)

	err := rpc.Register(master)

	if err != nil {

		log.Fatal("Format of service Task isn't correct. ", err)

	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", "4040")

	if err == nil {
		log.Fatal("Listen error: ", err)
	}

	fmt.Println("Serving RPC server on port 4040..")

	error := http.Serve(listener, nil)

	if error == nil {

		log.Fatal("Error serving: ", err)
	}

}
