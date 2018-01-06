package main

//-----------------------------------------------------------------------------
// Package factored import statement:
//-----------------------------------------------------------------------------

import (

	// Stdlib:
	"fmt"
	"os"

	// Community:
	"gopkg.in/alecthomas/kingpin.v2"
)

//-----------------------------------------------------------------------------
// Typedefs:
//-----------------------------------------------------------------------------

type stateData struct {
	name string
	fips int
	data []byte
}

type states map[string]stateData

//-----------------------------------------------------------------------------
// worker:
//-----------------------------------------------------------------------------

func worker(id int, jobs <-chan string, results chan<- stateData) {
	for j := range jobs {
		results <- stateData{name: j}
	}
}

//-----------------------------------------------------------------------------
// min:
//-----------------------------------------------------------------------------

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//-----------------------------------------------------------------------------
// Entry point:
//-----------------------------------------------------------------------------

func main() {

	// Parse the command-line:
	kingpin.MustParse(app.Parse(os.Args[1:]))

	// Variables:
	jobs := make(chan string, 50)
	results := make(chan stateData, 50)
	list := states{}

	// Initialize data:
	for _, state := range *argStates {
		list[state] = stateData{}
	}

	// Launch the workers:
	for i := 0; i < min(len(list), *flgMaxWorkers); i++ {
		go worker(i, jobs, results)
	}

	// Add jobs to the queue:
	for state := range list {
		jobs <- state
	}
	close(jobs)

	// Wait for the results:
	for i := 0; i < len(list); i++ {
		res := <-results
		list[res.name] = res
	}

	// Do something with the results:
	for k, v := range list {
		fmt.Println(k, v)
	}
}
