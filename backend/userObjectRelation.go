package main

import (
	"time"
)

// UserObjectRelation relates two objs
type UserObjectRelation struct {
	ID         int
	UserID     int
	ObjectUUID ObjectUUID // string, shortcut to make things faster
	Time       time.Time
	Purchased  bool
}

// CreateUserObjectRelation and add it to the database
func (db *Database) CreateUserObjectRelation(uor *UserObjectRelation) error {
	return db.userObjectRelationMap.Insert(uor)
}

func (db *Database) GetUORByID(id int) (*UserObjectRelation, error) {
	uor := new(UserObjectRelation)
	err := db.userObjectRelationMap.Get(uor, id)

	return uor, err
}

// UserObjectRelationByUserID that haven't been purchased
func (db *Database) UserObjectRelationByUserID(id int) (
	[]UserObjectRelation, error) {

	var uor []UserObjectRelation

	err := db.userObjectRelationMap.Select(&uor,
		"SELECT objectuuid FROM user_object_relation WHERE userid = ? AND purchased = ?",
		id, false)

	return uor, err
}
