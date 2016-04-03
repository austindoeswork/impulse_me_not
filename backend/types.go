package main

import ()

// types contains various types used by amazon

// UserUUID is unique user identifier
type UserUUID uint64

// ObjectUUID is an array of chars.
// TODO: be __very__ careful about how we parse and use strings
// len of this is 10
type ObjectUUID string
