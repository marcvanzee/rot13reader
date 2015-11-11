package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	bytesRead := 0
	b2 := make([]byte, 8)
	for {
		n, err := rot.r.Read(b2)
		if err == io.EOF {
			return bytesRead, io.EOF
		}
		b2 = encode(b2)
		copy(b[bytesRead:], b2)
		bytesRead += n
	}
}

func encode(str []byte) []byte {

	str2 := make([]byte, len(str))

	const a, z, A, Z = 'a', 'z', 'A', 'Z'

	for i, c := range str {
		value := c

		if c >= a && c <= z {
			value = (c-a+13)%26 + a
		} else if c >= A && c <= Z {
			value = (c-A+13)%26 + A
		}

		str2[i] = byte(value)
	}

	return str2
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
