package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUORByUserId(t *testing.T) {
	db := openTestDB()
	defer db.Close()

	user1 := User{UUID: getTestUserUUID()}
	user2 := User{UUID: getTestUserUUID()}

	db.CreateUser(&user1)
	db.CreateUser(&user2)

	object1 := Object{UUID: getTestObjectUUID()}
	object2 := Object{UUID: getTestObjectUUID()}

	uor1 := UserObjectRelation{UserID: 1, ObjectUUID: object1.UUID}
	uor2 := UserObjectRelation{UserID: 1, ObjectUUID: object2.UUID}
	uor3 := UserObjectRelation{UserID: 2, ObjectUUID: object1.UUID}

	db.CreateUserObjectRelation(&uor1)
	db.CreateUserObjectRelation(&uor2)
	db.CreateUserObjectRelation(&uor3)

	uors, err := db.UserObjectRelationByUserID(1)

	assert.Equal(t, nil, err)
	assert.NotNil(t, uors)
	assert.Len(t, uors, 2)

	//assert.Equal(t, uor1.ObjectUUID, uors[0].ObjectUUID)
	//assert.Equal(t, uor2.ObjectUUID, uors[1].ObjectUUID)
}

func TestGetUORById(t *testing.T) {
	db := openTestDB()
	defer db.Close()

	user1 := User{UUID: getTestUserUUID()}
	user2 := User{UUID: getTestUserUUID()}

	db.CreateUser(&user1)
	db.CreateUser(&user2)

	object1 := Object{UUID: getTestObjectUUID()}
	object2 := Object{UUID: getTestObjectUUID()}

	uor1 := UserObjectRelation{UserID: 1, ObjectUUID: object1.UUID}
	uor2 := UserObjectRelation{UserID: 1, ObjectUUID: object2.UUID}
	uor3 := UserObjectRelation{UserID: 2, ObjectUUID: object1.UUID}

	err := db.CreateUserObjectRelation(&uor1)
	assert.Nil(t, err)

	err = db.CreateUserObjectRelation(&uor2)
	assert.Nil(t, err)

	err = db.CreateUserObjectRelation(&uor3)
	assert.Nil(t, err)

	res1, err := db.GetUORByID(1)
	assert.Nil(t, err)
	assert.Equal(t, res1.UserID, 1)
}
