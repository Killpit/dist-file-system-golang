package main

import (
	"bytes"
	"fmt"
	"io"
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
	if pathKey.Filename != expectedPathName {
		t.Error(t, "have %s want %", pathKey.Filename, expectedOriginalKey)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts {
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "momsspecials"
	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(r)

	fmt.Println(string(b))

	if string(b) != string(data) {
		t.Errorf("want %s have %s", data, b)
	}

}