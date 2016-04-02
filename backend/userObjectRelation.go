package main

import (
	"time"
)

// UserObjectRelation relates two objs
type UserObjectRelation struct {
	ID        int
	ObjectID  int
	UserID    int
	AddTime   time.Time
	Purchased bool
}

// CreateUserObjectRelation and add it to the database
func (database *Database) CreateUserObjectRelation(uor *UserObjectRelation) error {
	return database.userObjectRelationMap.Insert(uor)
}
