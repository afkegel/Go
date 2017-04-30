package tutorial

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestRot13Reader(t *testing.T) {
	s1 := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{r: s1}

	b1 := make([]byte, 64)
	if _, err := r.Read(b1); err != io.EOF && err != nil {
		t.Error(err)
	}

	s2 := strings.NewReader("you cracked the code!")

	b2 := make([]byte, 64)
	if _, err := s2.Read(b2); err != io.EOF && err != nil {
		t.Error(err)
	}

	if string(b1) != string(b2) {
		t.Errorf("Rot13 decryption failed.")
	}
}

func TestRot13Reader_StringError(t *testing.T) {
	s := strings.NewReader(*new(string))
	r := Rot13Reader{r: s}

	b := make([]byte, 64)
	if _, err := r.Read(b); fmt.Sprint(err) != "empty string provided" {
		t.Error(err)
	}
}

func TestRot13Reader_SliceError(t *testing.T) {
	s := strings.NewReader("Slice test")
	r := Rot13Reader{r: s}

	b := make([]byte, 0)
	if _, err := r.Read(b); fmt.Sprint(err) != "slice has zero capacity" {
		t.Error(err)
	}
}
