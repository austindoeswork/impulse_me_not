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

	result2, err := db.UserByID(user2.UUID)

	assert.Equal(t, err, nil)
	assert.Equal(t, result2.UUID, user2.UUID)
}

func TestDuplicateUUID(t *testing.T) {
	db := openTestDB()
	defer db.Close()

	user1 := User{UUID: 1}
	user2 := User{UUID: 1}

	err := db.CreateUser(&user1)
	assert.Nil(t, err)

	err = db.CreateUser(&user2)
	assert.NotNil(t, err)
}

func TestUserByUUID(t *testing.T) {
	db := openTestDB()
	defer db.Close()

	user1 := User{UUID: 1}
	user2 := User{UUID: 2}

	db.CreateUser(&user1)
	db.CreateUser(&user2)

	result2, err := db.UserByUUID(user2.UUID)
	assert.Nil(t, err)
	assert.Equal(t, result2, &user2)

	// find bad uuid
	_, err = db.UserByUUID(0)
	assert.NotNil(t, err)
}
