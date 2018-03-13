package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"
)

type PoliceAPI []struct {
	District    string `json:"district,omitempty"`
	IncDatetime string `json:"inc_datetime"`
	IncNo       string `json:"inc_no"`
	Lcr         string `json:"lcr"`
	LcrDesc     string `json:"lcr_desc"`
	Location    struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"location,omitempty"`
}

func main() {
	response, err := http.Get("https://data.raleighnc.gov/resource/3bhm-we7a.json")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	
	var record PoliceAPI
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	
	fmt.Println(PolieAPI.District)
	
}
