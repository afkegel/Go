package tutorial

import (
	"bytes"
	"errors"
	"io"
	"strings"
)

// Rot13Reader provides methods to decode a rot13 encoded string to io.Reader
type Rot13Reader struct {
	r   io.Reader
	err error
}

// String method for a Rot13Reader returning the undecoded string.
func (r13 *Rot13Reader) String() string {
	buf := new(bytes.Buffer)
	_, r13.err = buf.ReadFrom(r13.r)

	return buf.String()
}

// Decode method for a Rot13Reader returning the decoded string.
func (r13 *Rot13Reader) Read(p []byte) (i int, err error) {

	if cap(p) == 0 {
		r13.err = errors.New("slice has zero capacity")
		return 0, r13.err
	}

	re := r13.decode()

	if re.err != nil {
		return 0, re.err
	}

	return re.r.Read(p)
	// this read method comes from the provided reader (e.g., strings.Reader)
}

func (r13 *Rot13Reader) decode() Rot13Reader {
	normal := "abcdefghijklmnopqrstuvwxyz.!? "
	rot13 := "nopqrstuvwxyzabcdefghijklm.!? "

	s1 := r13.String()
	s2 := ""

	if string(s1) == "" || string(s1) == "<nil>" {
		r13.err = errors.New("empty string provided")
		return *r13
	}

	for i := range s1 {
		idx := strings.Index(rot13, strings.ToLower(string(s1[i])))
		s2 = s2 + string(normal[idx])
	}
	return Rot13Reader{r: strings.NewReader(s2)}
}
