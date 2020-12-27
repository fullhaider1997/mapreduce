package main

import (
	Utility "MapReduce/Utility"
)

func main() {

	number_chunks := 30
	filename := "C:/Go/src/MapReduce/Book/pg-being_ernest.txt"

	Utility.Split_Data(filename, number_chunks)

}
