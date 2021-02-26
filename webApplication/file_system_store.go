package main

import (
	"encoding/json"
	"io"
)

//FileSystemPlayerStore implements PlayerStore using file system.
type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league   League
}

//NewFileSystemPlayerStore do the initial logic for FileSystemPlayerStore
func NewFileSystemPlayerStore(db io.ReadWriteSeeker) *FileSystemPlayerStore {
	db.Seek(0, 0)
	league, _ := NewLeague(db)
	return &FileSystemPlayerStore{
		database: db,
		league:   league,
	}
}

//GetLeague returns the players info in league
func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

//GetPlayerScore returns the players score
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

//RecordWin used to save the win count based on the player's name
func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}
