package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

//var testUUIDRandSrc = rand.NewSource(time.Now().UnixNano())
var testUUIDRandSrc = rand.NewSource(42)

// from stack overflow
// this is good for >1m calls
func getTestObjectUUID() ObjectUUID {
	n := 12
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, testUUIDRandSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = testUUIDRandSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return ObjectUUID(b)
}

// this is good for > 1m calls
func getTestUserUUID() UserUUID {
	//var id uint64
	return UserUUID(testUUIDRandSrc.Int63())

	//return UserUUID(id)
}

func TestObjectUUID(t *testing.T) {
	t.Skip("this test takes a little while, so disable")
	testUUIDRandSrc = rand.NewSource(time.Now().UnixNano())
	found := make(map[ObjectUUID]struct{})

	for i := 0; i < 1000000; i++ {
		val := getTestObjectUUID()
		if _, ok := found[val]; ok {
			t.Errorf("found repeate at %d!\n", i)
			break
		}
		found[val] = struct{}{}
	}
}

func TestUserUUID(t *testing.T) {
	t.Skip("this test takes a little while, so disable")
	testUUIDRandSrc = rand.NewSource(time.Now().UnixNano())
	found := make(map[UserUUID]struct{})

	for i := 0; i < 10000000; i++ {
		val := getTestUserUUID()
		if _, ok := found[val]; ok {
			t.Errorf("found repeate at %d!\n", i)
			break
		}
		found[val] = struct{}{}
	}
}

func TestCastString(t *testing.T) {
	goalFunc := func(a string) {
		t.Log(a)
	}

	uuid := getTestObjectUUID()
	fmt.Println("uuid:", uuid)

	//var defString string
	//defString = "asdfasdf"
	//def := reflect.ValueOf(defString)
	//fmt.Println("def:", def.Kind())

	rt := reflect.TypeOf(uuid)
	rv := reflect.ValueOf(uuid)
	fmt.Println("rt:", rt.Kind())
	fmt.Println("rv:", rv.CanAddr())
	//if rt.ConvertibleTo(string) {
	//rv.Convert(string)
	//}

	goalFunc(string(uuid))
}
