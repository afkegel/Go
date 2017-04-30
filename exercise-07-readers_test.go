package tutorial

import (
	"errors"
	"fmt"
	"io"
	"testing"
)

func TestMyReader(t *testing.T) {
	r := new(MyReader)

	// test correct functioning of myreader
	b := make([]byte, 1024, 2048)
	if ok, _ := Validate(r, b); !ok {
		t.Error("myreader failed")
	}

	// test zero capacity error
	b = make([]byte, 0)
	if _, err := Validate(r, b); fmt.Sprint(err) != "provided slice has 0 capacity" {
		t.Error("check for zero slice capacity failed")
	}

	// test byte other than 'A' error
	r.s = "B"
	b = make([]byte, 64)
	if _, err := Validate(r, b); fmt.Sprint(err) != "got byte other than 'A', want 'A'" {
		t.Error("check for correct byte read failed")
	}

	// test read zero bytes error
	r.s = "<nil>"
	if _, err := Validate(r, b); fmt.Sprint(err) != "read zero bytes" {
		t.Error("check for zero bytes read failed")
	}
}

// Validate function for an io.Reader copied from
// https://github.com/golang/tour/blob/master/reader/validate.go#L13 and
// adapted for testing.
func Validate(r io.Reader, b []byte) (bool, error) {
	i, o := 0, 0
	for ; i < 1<<20 && o < 1<<20; i++ { // test 1mb
		n, err := r.Read(b)
		for _, v := range b[:n] {
			if v != 'A' {
				return false, errors.New("got byte other than 'A', want 'A'")
			}
		}
		o += n
		if err != nil {
			return false, err
		}
	}
	if o == 0 {
		return false, errors.New("read zero bytes")
	}
	return true, nil
}
