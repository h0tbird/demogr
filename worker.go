package main

//-----------------------------------------------------------------------------
// Package factored import statement:
//-----------------------------------------------------------------------------

import (

	// Stdlib:
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

//-----------------------------------------------------------------------------
// Constants:
//-----------------------------------------------------------------------------

const apiBase = "https://www.broadbandmap.gov/broadbandmap"

//-----------------------------------------------------------------------------
// getStateFIPS:
//-----------------------------------------------------------------------------

func getStateFIPS(state string) (int, error) {

	// Send the request:
	resp, err := http.Get(apiBase + "/census/state/" + state + "?format=json")
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	// Read the response body:
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	// Decode the data:
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		return -1, err
	}

	// Extract the FIPS id:
	id := dat["Results"].(map[string]interface{})["state"].([]interface{})[0].(map[string]interface{})["fips"].(string)
	fips, err := strconv.Atoi(id)
	if err != nil {
		return -1, err
	}

	// Return:
	return fips, nil
}

//-----------------------------------------------------------------------------
// worker:
//-----------------------------------------------------------------------------

func worker(id int, jobs <-chan string, results chan<- stateData) {

	// Variables:
	var err error
	data := stateData{}

	// Job by job:
	for state := range jobs {

		// Get the state's FIPS:
		data.name = state
		data.fips, err = getStateFIPS(state)
		if err != nil {
			panic(err)
		}

		// Return:
		results <- data
	}
}
