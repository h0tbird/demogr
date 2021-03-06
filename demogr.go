package main

//-----------------------------------------------------------------------------
// Package factored import statement:
//-----------------------------------------------------------------------------

import (

	// Stdlib:
	"fmt"
	"os"
	"sort"

	// Community:
	"gopkg.in/alecthomas/kingpin.v2"
)

//-----------------------------------------------------------------------------
// Typedefs:
//-----------------------------------------------------------------------------

type demographic struct {
	population         float64
	households         float64
	incomeBelowPoverty float64
	medianIncome       float64
}

type stateData struct {
	name string
	fips string
	demographic
}

type states map[string]stateData

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

	// Output:
	switch *flgFormat {
	case "CSV":
		outputCSV(list)
	case "averages":
		outputAVG(list)
	}
}

//-----------------------------------------------------------------------------
// outputCSV:
//-----------------------------------------------------------------------------

func outputCSV(list states) {

	// Keep track of keys in an array:
	sortedKeys := make([]string, 0, len(list))
	for k := range list {
		sortedKeys = append(sortedKeys, k)
	}

	// Sort and print to stdout:
	sort.Strings(sortedKeys)
	for _, v := range sortedKeys {
		fmt.Printf("%s,%d,%d,%f,%f\n",
			list[v].name,               // string
			uint(list[v].population),   // uint
			uint(list[v].households),   // uint
			list[v].incomeBelowPoverty, // float64
			list[v].medianIncome)       // float64
	}
}

//-----------------------------------------------------------------------------
// outputAVG:
//-----------------------------------------------------------------------------

func outputAVG(list states) {

	// Variables:
	var numerator float64
	var denominator float64

	// Do the math:
	for _, v := range list {
		numerator = numerator + v.population*v.incomeBelowPoverty
		denominator = denominator + v.population
	}

	// Print to stdout:
	fmt.Println(numerator / denominator)
}
