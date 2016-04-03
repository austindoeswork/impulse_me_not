package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestGetUORByUserId(t *testing.T) {
	db := openTestDB()
	defer db.Close()

	db.userMap.TraceOn("modl:", log.New(os.Stdout, "", log.Lshortfile))
	defer db.userMap.TraceOff()

	user1 := User{ID: 1}
	user2 := User{ID: 2}

	object1 := Object{UUID: getTestObjectUUID()}
	object2 := Object{UUID: getTestObjectUUID()}

	uor1 := UserObjectRelation{UserID: user1.ID, ObjectUUID: object1.UUID}
	uor2 := UserObjectRelation{UserID: user1.ID, ObjectUUID: object2.UUID}
	uor3 := UserObjectRelation{UserID: user2.ID, ObjectUUID: object1.UUID}

	err := db.CreateUserObjectRelation(&uor1)
	assert.Equal(t, nil, err)
	err = db.CreateUserObjectRelation(&uor2)
	assert.Equal(t, nil, err)
	err = db.CreateUserObjectRelation(&uor3)
	assert.Equal(t, nil, err)

	uors, err := db.UserObjectRelationByUserID(user1.ID)

	assert.Equal(t, nil, err)
	assert.NotNil(t, uors)
	assert.Len(t, uors, 2)

	assert.Equal(t, uor1.ObjectUUID, uors[0].ObjectUUID)
	assert.Equal(t, uor2.ObjectUUID, uors[1].ObjectUUID)

}

func TestGetUORById(t *testing.T) {
	db := openTestDB()
	defer db.Close()

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
	assert.Equal(t, 1, res1.UserID)
}
