package main

import (
	"fmt"
)

// User define
type User struct {
	ID   int
	UUID UserUUID // UNSIGNED BIG INT, uint64
}

// CreateUser and insert it into the db
func (db *Database) CreateUser(user *User) error {
	return db.userMap.Insert(user)
}

// UserByID from the DB
func (db *Database) UserByID(id int) (*User, error) {
	user := new(User)

	err := db.userMap.Get(user, id)
	// TODO: real error checking
	return user, err
}

// UserByUUID gets a user by uuid
func (db *Database) UserByUUID(uuid UserUUID) (*User, error) {
	var users []User
	err := db.userMap.Select(&users,
		"SELECT * FROM user WHERE uuid = ?", uuid)

	if err != nil {
		return nil, err
	} else if len(users) != 1 {
		return nil, fmt.Errorf("bad num users found, expected 1 got %d\n",
			len(users))
	}

	// get value so the array will be free'd
	// the gc will make this work
	user := users[0]
	return &user, nil
}
