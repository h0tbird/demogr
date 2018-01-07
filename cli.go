package main

//-----------------------------------------------------------------------------
// Package factored import statement:
//-----------------------------------------------------------------------------

import (

	// Stdlib:
	"fmt"
	"sort"
	"strings"

	// Community:
	"gopkg.in/alecthomas/kingpin.v2"
)

//-----------------------------------------------------------------------------
// Listings:
//-----------------------------------------------------------------------------

var (

	// U.S. states:
	statesList = []string{
		"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado",
		"Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho",
		"Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine",
		"Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi",
		"Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey",
		"New Mexico", "New York", "North Carolina", "North Dakota", "Ohio",
		"Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina",
		"South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia",
		"Washington", "West Virginia", "Wisconsin", "Wyoming"}

	// Data versions:
	dataVersions = []string{
		"jun2011", "dec2011", "jun2012", "dec2012", "jun2013", "dec2013", "jun2014"}
)

//-----------------------------------------------------------------------------
// Command, flags and arguments:
//-----------------------------------------------------------------------------

var (

	// Command:
	app = kingpin.New("demogr", "Retrieves demographic data for a specified set of U.S. states from a public API and outputs that data in the requested format.")

	// Flags:
	flgFormat = app.Flag("format", "Output-format parameter [CSV|averages]").
			Default("CSV").Enum("CSV", "averages")

	flgMaxWorkers = app.Flag("max-workers", "Maximum number of concurrent workers.").
			Default("5").Int()

	flgDataVersion = app.Flag("data-version", "Specify the data version.").
			Default("jun2014").Enum(dataVersions...)

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
