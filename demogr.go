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
// Fixed data:
//-----------------------------------------------------------------------------

var states = []string{"Alabama", "Alaska", "Arizona", "Arkansas",
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
	flgFormat = app.Flag("format",
		"Output-format parameter [ CSV | averages ].").
		Default("CSV").HintOptions("CSV", "averages").String()

	// Arguments:
	argStates = statesCSV(app.Arg("states",
		"Comma delimited list of U.S. states.").
		Required(), states)
)

//-----------------------------------------------------------------------------
// Custom CSV states parser:
//-----------------------------------------------------------------------------

type stateSlice []string

func (s *stateSlice) Set(value string) error {
	*s = strings.Split(value, ",")
	for _, state := range *s {
		i := sort.SearchStrings(states, state)
		if i == len(states) || states[i] != state {
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
// Entry point:
//-----------------------------------------------------------------------------

func main() {

	// Parse the command-line:
	kingpin.MustParse(app.Parse(os.Args[1:]))

	// Initialize the data:
	for _, state := range *argStates {
		fmt.Println(state)
	}
}
