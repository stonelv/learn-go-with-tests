package main

import (
	"encoding/json"
	"fmt"
	"io"
)

//League represents the players info
type League []Player

//Find returns the Player Object by name
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

//NewLeague read players info from io.Reader and return them.
func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return league, err
}
