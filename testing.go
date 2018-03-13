package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"
)

type Entry struct {
	Incident	[]string	`json:="inc_no"`
}

func main() {
	response, err := http.Get("https://data.raleighnc.gov/resource/3bhm-we7a.json")
	i := 0
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	var entryObject Entry
	
	json.Unmarshal(responseData, &entryObject)
	fmt.Println(entryObject.Incident[i])
	fmt.Println(newline)
}
