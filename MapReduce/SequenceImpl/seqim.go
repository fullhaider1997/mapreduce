package main

//
// simple sequential MapReduce.
//
// go run mrsequential.go ../mrapps/wc.so pg*.txt
//

import (
	mr "MapReduce/Worker"
	wc "MapReduce/wrapper"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

// for sorting by key.
type ByKey []mr.KeyValue

// for sorting by key.
func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

func main() {

	filename := "C:/Go/src/MapReduce/Book/pg-being_ernest.txt"
	//
	// read each input file,
	// pass it to Map,
	// accumulate the intermediate Map output.
	//
	intermediate := []mr.KeyValue{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("cannot open %v", filename)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read %v", filename)
	}
	file.Close()

	kva := wc.Mapf(filename, string(content))
	intermediate = append(intermediate, kva...)

	//
	// a big difference from real MapReduce is that all the
	// intermediate data is in one place, intermediate[],
	// rather than being partitioned into NxM buckets.
	//

	/*

		Before the reduce, we have sort the array of keys and values
		by keys so that occurences of the same key are grouped together.


	*/
	sort.Sort(ByKey(intermediate))

	oname := "mr-out-0"
	ofile, _ := os.Create(oname)

	//
	// call Reduce on each distinct key in intermediate[],
	// and print the result to mr-out-0.
	//

	i := 0
	for i < len(intermediate) {
		j := i + 1
		for j < len(intermediate) && intermediate[j].Key == intermediate[i].Key {
			j++
		}

		values := []string{}

		for k := i; k < j; k++ {
			values = append(values, intermediate[k].Value)
		}

		output := wc.Reducef(intermediate[i].Key, values)

		// this is the correct format for each line of Reduce output.
		fmt.Fprintf(ofile, "%v %v\n", intermediate[i].Key, output)

		i = j
	}

	ofile.Close()
}
