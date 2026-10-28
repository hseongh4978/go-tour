package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13Char(b byte) byte {
	switch {
	case 'a' <= b && b <= 'z':
		return 'a' + (b-'a'+13)%26
	case 'A' <= b && b <= 'Z':
		return 'A' + (b-'A'+13)%26
	default:
		return b
	}
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := range b {
		b[i] = rot13Char(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
