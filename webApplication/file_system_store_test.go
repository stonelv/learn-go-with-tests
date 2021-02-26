package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Lyle", "Wins": 10},
			{"Name": "Josh", "Wins": 20}
		]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetLeague()
		want := []Player{
			{"Lyle", 10},
			{"Josh", 20},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Lyle", "Wins": 10},
			{"Name": "Josh", "Wins": 22}
		]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Josh")
		want := 22

		assertScoreEquals(got, want, t)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Lyle", "Wins": 10},
			{"Name": "Josh", "Wins": 22}
		]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Josh")

		got := store.GetPlayerScore("Josh")
		want := 23
		assertScoreEquals(got, want, t)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Lyle", "Wins": 10},
			{"Name": "Josh", "Wins": 22}
		]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(got, want, t)
	})
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tempFile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tempFile.Write([]byte(initialData))

	removeFile := func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}

	return tempFile, removeFile
}

func assertScoreEquals(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
