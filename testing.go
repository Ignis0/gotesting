package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"
)

type Response struct {
	Pokemon []Pokemon
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo string	`json:"inc_no"`
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

	fmt.Println(responseData)
	
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	
	fmt.Println(len(responseObject.Pokemon))

	//for i := 0; i < len(responseObject.Pokemon); i++ {
	//	fmt.Println(responseObject.Pokemon[i])
	//}
}
