package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"
)

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
	
	fmt.Println(string(responseData))

	entry []string `json: "inc_no"`
	
	for i := 0; i < len(entry); i++ {
		fmt.Println(entry[i])
	}
	fmt.Println("done")
	
}
