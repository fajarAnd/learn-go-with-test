package main

import (
	"github.com/fajarAnd/learn-go-with-test/http-server/apps"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(apps.PlayerServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}