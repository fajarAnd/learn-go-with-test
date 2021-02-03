package main

import (
	"github.com/fajarAnd/learn-go-with-test/http-server/apps"
	"log"
	"net/http"
)

//main.go
type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	//server := &apps.PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5000", &apps.PlayerServer{}); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
