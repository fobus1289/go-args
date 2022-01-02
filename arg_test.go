package go_args

import (
	"testing"
)

var keys = []string{"port", "address"}

func TestSetErrorTrueArgs(t *testing.T) {
	if _, err := GetMapArgs(true, keys...); err != nil {
		t.Log("success")
	} else {
		t.Error("Error")
	}
}

func TestSetErrorFalseArgs(t *testing.T) {
	if _, err := GetMapArgs(false, keys...); err != nil {
		t.Error("Error")
	} else {
		t.Log("success")
	}
}
