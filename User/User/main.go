package main

import (
	Utility "MapReduce/Utility"
	"log"
	"os"
)

func main() {

	filename := "C:/Go/src/MapReduce/Book/pg-being_ernest.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("cannot open %v", filename)
	}

	//content, err := ioutil.ReadAll(file)
	//if err != nil {
	//	log.Fatalf("cannot read %v", filename)
	//}

	Utility.Split_Data()

	file.Close()

}
