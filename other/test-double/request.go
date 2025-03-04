package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type people struct {
	Number int `json: "number"`
}

type GetWebRequest interface {
	FetchBytes(url string) []byte
}

type LiveGetWebRequest struct {
}

func (LiveGetWebRequest) FetchBytes(url string) []byte {
	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)

	//if body != nil{
	//	defer body.Close()
	//}

	if readErr != nil {
		log.Fatal(readErr)
	}

	log.Print(body, "badan")

	return body

}
