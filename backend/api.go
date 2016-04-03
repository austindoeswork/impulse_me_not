package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	//"strconv"
)

// HandleTest is a test handler
func HandleTest(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("hey, whats up"))
}

// HandleNewUser when the app is first downloaded.
// The form of this request is /newUser?<uuid>.
// If the user is not new, then there open impulses are returned.
// Otherwise a new user is created.
func HandleNewUser(w http.ResponseWriter, r *http.Request) {
}

// HandleOpenImpulses returns all the user's open impulses.
// The form of this request is <server>/openImpulses/<uuid>.
// Puts a json array of uuids into the object
func HandleOpenImpulses(w http.ResponseWriter, request *http.Request) {
	/*
		rawQuery := request.URL.RawQuery

		uuid, err := strconv.ParseUint(rawQuery, 10, 64)
		if err != nil {
			w.Write([]byte("bad uuid"))
			return
		}

		user, err := db.UserByUUID(UserUUID(uuid))
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
		arr := make([]ObjectUUID, len(uors))
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
	*/
}

// HandleAddImpulse adds an impulse for a user.
// Endpoint is /addItem.
// body should be {user = user_uuid, object = object_uuid}
func HandleAddImpulse(w http.ResponseWriter, r *http.Request) {
	js, err := ioutil.ReadAll(r.Body)
	if err != nil || js == nil || len(js) == 0 {
		w.Write([]byte("bad body"))
		return
	}

	var body map[string]interface{}
	err = json.Unmarshal(js, &body)

	// TODO: actuall use values instead of _
	_, hasUser := body["user"]
	_, hasObject := body["object"]

	if err != nil || !hasUser || !hasObject {
		w.Write([]byte("bad json"))
		return
	}
}
