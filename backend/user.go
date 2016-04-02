package main

// User define
type User struct {
	ID   int
	UUID int
}

// CreateUser and insert it into the db
func (db *Database) CreateUser(user *User) error {
	return db.userMap.Insert(user)
}

// GetUserByID from the DB
func (db *Database) GetUserByID(id int) (*User, error) {
	user := new(User)

	err := db.userMap.Get(user, id)
	// TODO: real error checking
	return user, err
}
