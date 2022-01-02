package go_args

import (
	"testing"
)

var keys = []string{"port", "address"}

func TestSetErrorTrueArgs(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("success")
		} else {
			t.Error("Error")
		}
	}()

	GetMapArgs(true, keys...)

}

func TestSetErrorFalseArgs(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		} else {
			t.Log("success")
		}
	}()

	GetMapArgs(false, keys...)
}
