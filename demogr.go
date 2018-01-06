package main

//-----------------------------------------------------------------------------
// Package factored import statement:
//-----------------------------------------------------------------------------

import (

	// Stdlib:
	"fmt"
	"os"
	"sort"
	"strings"

	// Community:
	"gopkg.in/alecthomas/kingpin.v2"
)

//-----------------------------------------------------------------------------
// US states list:
//-----------------------------------------------------------------------------

var statesList = []string{"Alabama", "Alaska", "Arizona", "Arkansas",
	"California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia",
	"Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky",
	"Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota",
	"Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire",
	"New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota",
	"Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island",
	"South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont",
	"Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}

//-----------------------------------------------------------------------------
// Command, flags and arguments:
//-----------------------------------------------------------------------------

var (

	// Command:
	app = kingpin.New("demogr", "Retrieves demographic data for a specified set of U.S. states from a public API and outputs that data in the requested format.")

	// Flags:
	flgFormat = app.Flag("format", "Output-format parameter [CSV|averages]").
			Default("CSV").HintOptions("CSV", "averages").String()

	flgMaxWorkers = app.Flag("max-workers", "Maximum number of concurrent workers.").
			Default("5").Int()

	// Arguments:
	argStates = statesCSV(app.Arg("states", "Comma delimited list of U.S. states.").
			Required(), statesList)
)

//-----------------------------------------------------------------------------
// Custom CSV states parser:
//-----------------------------------------------------------------------------

type stateSlice []string

func (s *stateSlice) Set(value string) error {
	*s = strings.Split(value, ",")
	for _, state := range *s {
		i := sort.SearchStrings(statesList, state)
		if i == len(statesList) || statesList[i] != state {
			return fmt.Errorf("'%s' is not a valid US state", state)
		}
	}
	return nil
}

func (s *stateSlice) String() string {
	return ""
}

func statesCSV(s kingpin.Settings, states []string) *stateSlice {
	target := &stateSlice{}
	s.SetValue(target)
	return target
}

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
