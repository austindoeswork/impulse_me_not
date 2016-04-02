package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleHandle(t *testing.T) {
	recorder := getHTTPResponse("/test")
	assert.Equal(t, recorder.Body.String(), "hey, whats up")
}
