package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	size := len(b)
	for i := 0; i < size; i++ {
		b[i] = 'A'
	}
	return size, nil
}

func main() {
	reader.Validate(MyReader{})
}
