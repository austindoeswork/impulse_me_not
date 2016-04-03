package main

// Object purchasable from amazon
type Object struct {
	// ID of the record
	ID int

	// UUID of the object, from amazon
	UUID int

	// AddCount is the number of times people have considered buying
	AddCount int

	// PurchaseCount is the number of times someone actually bought
	PurchaseCount int
}

// CreateObject and insert it into the database
func (database *Database) CreateObject(object *Object) error {
	return database.objectMap.Insert(object)
}

// ObjectsByUUID returns an object by its uuid
func (database *Database) ObjectsByUUID(uuid int) ([]Object, error) {
	return nil, nil
}
