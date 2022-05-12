package gocapcv

import (
	"testing"

	"gocv.io/x/gocv"
)

func Test1(t *testing.T) {
	mat, err := CaptureScreen(0, 0, 100, 100)
	if err != nil {
		t.Fatalf("%s", err)
	}

	res := gocv.IMWrite("test.png", mat)
	if !res {
		t.Fatalf("Failed to write image")
	}
}
