package tutorial

// MyReader provides a reader to an embedded string.
type MyReader struct {
	s string
}

type errorMyReader struct {
	s string
}

func (e *errorMyReader) Error() string {
	return e.s
}

// Read method to MyReader
func (m *MyReader) Read(b []byte) (n int, err error) {
	n = 0
	if cap(b) == 0 {
		return n, &errorMyReader{"provided slice has 0 capacity"}
	}
	if m.s == "<nil>" {
		return n, &errorMyReader{"read zero bytes"}
	}
	if m.s == "" {
		m.s = "A"
	}
	for n = range b {
		b[n] = byte([]byte(m.s)[0])
	}
	return n, nil
}
