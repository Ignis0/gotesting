package main


import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type passData []struct {
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
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	var responseObject passData
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(len(responseObject))
	
	for i := 0; i < len(responseObject); i++ {
		fmt.Println(responseObject[i].District, " ",
			    responseObject[i].IncDatetime, " ",
			    responseObject[i].IncNo, " ",
			    responseObject[i].Lcr, " ",
			    responseObject[i].LcrDesc, " ",
			    responseObject[i].Location.Type, " ",
			    responseObject[i].Location.Coordinates)
	}

}
