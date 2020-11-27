package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func GetAstronauts(getWebRequest GetWebRequest) int {
	url := "http://api.open-notify.org/astros.json"
	bodyBytes := getWebRequest.FetchBytes(url)
	peopleResult := people{}
	jsonErr := json.Unmarshal(bodyBytes, &peopleResult)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	log.Print("wow >>", peopleResult)

	return peopleResult.Number
}

func main() {
	liveClient := LiveGetWebRequest{}
	number := GetAstronauts(liveClient)

	fmt.Println(number)
}
