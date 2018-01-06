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
// url2json:
//-----------------------------------------------------------------------------

func url2json(url string) (map[string]interface{}, error) {

	// Variables:
	var dat map[string]interface{}

	// Send the request:
	resp, err := http.Get(url)
	if err != nil {
		return dat, err
	}
	defer resp.Body.Close()

	// Read the response body:
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dat, err
	}

	// Decode the data:
	if err := json.Unmarshal(body, &dat); err != nil {
		return dat, err
	}

	// Return:
	return dat, nil
}

//-----------------------------------------------------------------------------
// getStateFIPS:
//-----------------------------------------------------------------------------

func getStateFIPS(state string) (string, error) {

	// Send the request:
	dat, err := url2json(apiBase + "/census/state/" + state + "?format=json")
	if err != nil {
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
	dat, err := url2json(apiBase + "/demographic/" + *flgDataVersion + "/state/ids/" + fips + "?format=json")
	if err != nil {
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
