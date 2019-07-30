package main

import (
	"testing"
)

func TestConnect(T *testing.T) {
	_, err := connect()
	if err != nil {
		T.Fail()
	}
}
