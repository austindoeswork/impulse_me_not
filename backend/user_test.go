package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUser(t *testing.T) {
	db := openTestDB()
	defer db.Close()

	user1 := User{UUID: 1}
	user2 := User{UUID: 2}

	err := db.CreateUser(&user1)
	assert.Nil(t, err)

	err = db.CreateUser(&user2)
	assert.Nil(t, err)

	result2, err := db.GetUserByID(2)

	assert.Equal(t, err, nil)
	assert.Equal(t, result2.UUID, user2.UUID)
}

func TestBadUUID(t *testing.T) {
	db := openTestDB()
	defer db.Close()

	user1 := User{UUID: 1}
	user2 := User{UUID: 1}

	err := db.CreateUser(&user1)
	assert.Nil(t, err)

	err = db.CreateUser(&user2)
	assert.NotNil(t, err)
}
