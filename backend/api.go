package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// HandleTest is a test handler
func HandleTest(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("hey, whats up"))
}

// HandleOpenImpulses returns all the user's open impulses.
// The form of this request is <server>/openImpulses/<uuid>.
// Puts a json array of uuids into the object
func HandleOpenImpulses(w http.ResponseWriter, request *http.Request) {
	rawQuery := request.URL.RawQuery

	uuid, err := strconv.Atoi(rawQuery)
	if err != nil {
		w.Write([]byte("bad uuid"))
		return
	}

	user, err := db.UserByUUID(uuid)
	if err != nil {
		w.Write([]byte("bad uuid"))
		return
	}

	// get all open purchases
	uors, err := db.UserObjectRelationByUserID(user.ID)
	if err != nil {
		w.Write([]byte("problem getting objects"))
		return
	}

	// build json array for response
	arr := make([]int, len(uors))
	for i, uor := range uors {
		arr[i] = uor.ObjectUUID

		if err != nil {
			w.Write([]byte("problem encoding record!"))
			return
		}
	}

	js, err := json.Marshal(arr)
	if err != nil {
		w.Write([]byte("error marshalling js"))
	} else {
		w.Write(js)
	}

}
