package main

import (
	"database/sql/driver"
	"fmt"
)

// types contains various types used by amazon

// UserUUID is unique user identifier
type UserUUID uint64

// ObjectUUID is an array of chars.
// TODO: be __very__ careful about how we parse and use strings
// len of this is 10
type ObjectUUID string

// Value for converting to msql
func (uuid ObjectUUID) Value() (driver.Value, error) {
	fmt.Println("converted:", uuid)
	return string(uuid), nil
}

// Scan the value from msql
func (uuid *ObjectUUID) Scan(src interface{}) error {
	*uuid = ObjectUUID(string(src.([]uint8)))
	fmt.Println("scanned:", *uuid)
	return nil
}
