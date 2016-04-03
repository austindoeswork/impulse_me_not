package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestUserGet(t *testing.T) {
	db := openTestDB()

	db.userMap.TraceOn("modl:", log.New(os.Stdout, "", log.Lshortfile))
	defer db.Close()

	user1 := User{UUID: 1}

	uuid2 := getTestUserUUID()
	uuid2 = getTestUserUUID()
	uuid2 = getTestUserUUID()
	//uuid2 = UserUUID(2)
	user2 := User{UUID: uuid2}

	err := db.CreateUser(&user1)
	assert.Equal(t, nil, err)

	err = db.CreateUser(&user2)
	assert.Equal(t, nil, err)

	result2, err := db.UserByID(user2.ID)

	assert.Equal(t, nil, err)
	assert.NotNil(t, result2)
	assert.Equal(t, result2.UUID, user2.UUID)

	db.userMap.TraceOff()
}

func TestUserDuplicateUUID(t *testing.T) {
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
	db.userMap.TraceOn("modl:", log.New(os.Stdout, "", log.Lshortfile))
	defer db.Close()

	user1 := User{UUID: 1}

	uuid2 := getTestUserUUID()
	user2 := User{UUID: uuid2}

	err := db.CreateUser(&user1)
	assert.Equal(t, nil, err)

	err = db.CreateUser(&user2)
	assert.Equal(t, nil, err)

	result2, err := db.UserByUUID(user2.UUID)

	assert.Equal(t, nil, err)
	assert.NotNil(t, result2)
	assert.Equal(t, result2.UUID, user2.UUID)

	db.userMap.TraceOff()

}
