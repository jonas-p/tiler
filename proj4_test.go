package main

import (
	"testing"
)

func TestNewProj(t *testing.T) {
	sweref99 := "+proj=utm +zone=33 +ellps=GRS80 +towgs84=0,0,0,0,0,0,0 +units=m +no_defs"
	_, err := NewProj(sweref99)
	if err != nil {
		t.Error(err)
	}
	
	_, err = NewProj("this isnt correct")
	if err == nil {
		t.Error("err shouldn't be nil")
	}
}
