package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUORByUserId(t *testing.T) {
	db := openTestDB()
	defer db.Close()

	user1 := User{UUID: 3}
	user2 := User{UUID: 4}

	db.CreateUser(&user1)
	db.CreateUser(&user2)

	uor1 := UserObjectRelation{UserID: 1, ObjectUUID: 2}
	uor2 := UserObjectRelation{UserID: 1, ObjectUUID: 3}
	uor3 := UserObjectRelation{UserID: 2, ObjectUUID: 2}

	db.CreateUserObjectRelation(&uor1)
	db.CreateUserObjectRelation(&uor2)
	db.CreateUserObjectRelation(&uor3)

	uors, err := db.UserObjectRelationByUserID(1)

	assert.Equal(t, nil, err)
	assert.NotNil(t, uors)
	assert.Len(t, uors, 2)

	assert.Equal(t, uor1.ObjectUUID, uors[0].ObjectUUID)
	assert.Equal(t, uor2.ObjectUUID, uors[1].ObjectUUID)
}
