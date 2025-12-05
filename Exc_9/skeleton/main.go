package main

import (
	"exc9/mapred"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

// Main function
func main() {
	// read file
	_, file, _, ok := runtime.Caller(0) // gets information on the file of the exe, i.e. main.go
	if !ok {
		fmt.Println("Could not determine file path.")
		return
	}
	// Get the directory of the file
	dir := filepath.Dir(file) // gets the dir of the file, here .../skeleton

	path := filepath.Join(dir, "/res/meditations.txt") // txt file path 
	raw, err := os.ReadFile(path) // read as byte array 
	if err != nil{
		fmt.Printf("%s", err.Error())
		return
	}


	var text []string
	for line := range strings.Lines(string(raw)) { // convert bayte array to lines
		text = append(text, line)
	}

	var mr mapred.MapReduce
	results := mr.Run(text)

	// print your result to stdout

	// convert map to list
	sorted := []mapred.KeyValue{}
	for key, value := range results {
		sorted = append(sorted, mapred.KeyValue{Key: key, Value: value})
	}
	// so i con sort it
	sort.Slice(sorted, func(i, j int) bool {return sorted[j].Value < sorted[i].Value}) // descending
	
	// the txt has over 6000 lines, so limit output to only best 25
	for i := range sorted[:25]{
		fmt.Printf("%s -> %v\n", sorted[i].Key, sorted[i].Value)
	}

}
