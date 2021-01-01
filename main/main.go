package main

import (
	mapreduce "User-mapreduce/MapReduce/master-worker"
)

func main() {

	number_chunks := 3
	filename := "C:/Go/src/User-mapreduce/MapReduce/Book/pg-being_ernest.txt"

	mapreduce.Split_Data(filename, number_chunks)
	mapreduce.Start_Job(number_chunks)

}
