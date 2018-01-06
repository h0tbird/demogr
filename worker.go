package main

//-----------------------------------------------------------------------------
// Package factored import statement:
//-----------------------------------------------------------------------------

import (

	// Stdlib:
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//-----------------------------------------------------------------------------
// Constants:
//-----------------------------------------------------------------------------

const apiBase = "https://www.broadbandmap.gov/broadbandmap"

//-----------------------------------------------------------------------------
// getStateFIPS:
//-----------------------------------------------------------------------------

func getStateFIPS(state string) (string, error) {

	// Send the request:
	resp, err := http.Get(apiBase + "/census/state/" + state + "?format=json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body:
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Decode the data:
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		return "", err
	}

	// Extract the FIPS id:
	fips := dat["Results"].(map[string]interface{})["state"].([]interface{})[0].(map[string]interface{})["fips"].(string)
	return fips, nil
}

//-----------------------------------------------------------------------------
// getStateData:
//-----------------------------------------------------------------------------

func getStateData(fips string) (demographic, error) {

	// Variables:
	dmgr := demographic{}

	// Send the request:
	resp, err := http.Get(apiBase + "/demographic/jun2014/state/ids/" + fips + "?format=json")
	if err != nil {
		return dmgr, err
	}
	defer resp.Body.Close()

	// Read the response body:
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dmgr, err
	}

	// Decode the data:
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		return dmgr, err
	}

	// Extract population, households, incomeBelowPoverty and medianIncome:
	dmgr.population = dat["Results"].([]interface{})[0].(map[string]interface{})["population"].(float64)
	dmgr.households = dat["Results"].([]interface{})[0].(map[string]interface{})["households"].(float64)
	dmgr.incomeBelowPoverty = dat["Results"].([]interface{})[0].(map[string]interface{})["incomeBelowPoverty"].(float64)
	dmgr.medianIncome = dat["Results"].([]interface{})[0].(map[string]interface{})["medianIncome"].(float64)

	return dmgr, nil
}

//-----------------------------------------------------------------------------
// worker:
//-----------------------------------------------------------------------------

func worker(id int, jobs <-chan string, results chan<- stateData) {

	// Variables:
	var err error
	state := stateData{}

	// Job by job:
	for state.name = range jobs {

		// Get the state's FIPS:
		state.fips, err = getStateFIPS(state.name)
		if err != nil {
			panic(err)
		}

		// Get the state's data:
		state.demographic, err = getStateData(state.fips)
		if err != nil {
			panic(err)
		}

		// Return:
		results <- state
	}
}
