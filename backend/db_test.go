package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// main testing function to reset the db between runs
func openTestDB() *Database {
	clearDB()
	db, err := OpenDatabase(testDBPath)
	if err != nil {
		panic(err)
	}

	return db
}

const (
	testDBPath = "test.db"
)

func TestOpenNonexistantDB(t *testing.T) {
	_, err := OpenDatabase("badpath/bad.db")
	t.Log("db err is:", err)
	assert.NotNil(t, err)
}

// clears away the db (if any exists)
func clearDB() {
	os.Remove(testDBPath)
}
