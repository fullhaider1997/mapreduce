package Utility

import (
	master "MapReduce/Master"
	worker "MapReduce/Worker"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Split_Data(filename string, number_chunks int) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("cannot open %v", filename)
	}

	f, _ := file.Stat()

	//Calculate the size of each chunck
	chunckSize := int(f.Size()) / number_chunks

	//fmt.Println(chunckSize)

	var pos int64 = 0

	/*
	  Iterate through file based on number of chuncks and read chunkSize of bytes from the file.
	  After create new file and write new chunck of data into it.

	*/

	for i := 0; i < number_chunks; i++ {

		pos, _ = file.Seek(pos, 0)

		byteSlice := make([]byte, chunckSize)
		n, _ := file.Read(byteSlice)

		fmt.Println("Current position is ", pos)

		fmt.Println("Number of bytes read:", n)

		pos = pos + int64(chunckSize)

		f, err := os.Create("C:/Go/src/MapReduce/Chunks/chunk-" + strconv.Itoa(i) + ".txt")
		defer f.Close()

		if err == nil {
			fmt.Println("Error opening the file")
		}

		f.Write(byteSlice)

	}

	//Start number_chuncks workers
	master.Start()
	worker.Start_Worker(number_chunks)

}
