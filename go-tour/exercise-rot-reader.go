package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		if 'A' <= p[i] && p[i] <= 'Z' {
			p[i] = (p[i] + 13)
			if p[i] > 'Z' {
				p[i] -= 26
			}
		} else if 'a' <= p[i] && p[i] <= 'z' {
			p[i] = (p[i] + 13)
			if p[i] > 'z' {
				p[i] -= 26
			}
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
