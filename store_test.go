package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "//////d41d8cd98f00b204e9800998ecf8427e"
	expectedPathName := "//////d41d8cd98f00b204e9800998ecf8427e"
	
	if pathKey.PathName != expectedPathName {
		t.Error(t, "have %s want %", pathKey.PathName, expectedPathName)
	}
	if pathKey.Original != expectedPathName {
		t.Error(t, "have %s want %", pathKey.Original, expectedOriginalKey)
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