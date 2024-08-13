package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathname := CASPathTransformFunc(key)
	expectedPathName := "//////"
	fmt.Println(pathname)
	if pathname != expectedPathName {
		t.Error(t, "have %s want %", pathname, expectedPathName)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts {
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))
	s.writeStream("myspecialpicture", data)
	if err := s.writeStream("myspecialpicture", data); err != nil {
		t.Error(err)
	}
	
}