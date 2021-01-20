package main

import "testing"

type StubPlayerStore struct {
	score map[string]int
	winCalls []string
}

func TestLeague(t *testing.T) {
	//store := StubPlayerStore{}
	//server := &PlayerServer{&store}
}