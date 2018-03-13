package main

import (
	"fmt"
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
	
	url := "https://data.raleighnc.gov/resource/3bhm-we7a.json"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer response.Body.Close()
	
	var record PoliceAPI
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	
	fmt.Println(PoliceAPI.District)
	
}
