package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleHandle(t *testing.T) {
	recorder := getHTTPResponse("/test")
	assert.Equal(t, recorder.Body.String(), "hey, whats up")
}

func TestGetImpulses(t *testing.T) {
	/*
		db = openTestDB()
		defer db.Close()

		user1 := User{UUID: 10}
		user2 := User{UUID: 15}

		db.CreateUser(&user1)
		db.CreateUser(&user2)

		uor1 := UserObjectRelation{UserID: 1, ObjectUUID: 2}
		uor2 := UserObjectRelation{UserID: 1, ObjectUUID: 3}
		uor3 := UserObjectRelation{UserID: 2, ObjectUUID: 2}

		db.CreateUserObjectRelation(&uor1)
		db.CreateUserObjectRelation(&uor2)
		db.CreateUserObjectRelation(&uor3)

		uors, err := db.UserObjectRelationByUserID(1)

		assert.Nil(t, err)
		assert.NotNil(t, uors)
		assert.Len(t, uors, 2)

		recorder := getHTTPResponse("/openImpulses?10")
		assert.Equal(t, "[2,3]", recorder.Body.String())
	*/
}
