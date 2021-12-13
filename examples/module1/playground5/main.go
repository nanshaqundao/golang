package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func rot13substitution(b byte) byte {
	var s byte
	if (65 <= b) && (b <= 90) {
		s = b + 13
		if s > 90 {
			s = 65 + (s - 91)
		}
		return s
	}
	if (97 <= b) && (b <= 122) {
		s = b + 13
		if s > 122 {
			s = 97 + (s - 123)
		}
		return s
	}
	s = b
	return s

}

type rot13Reader struct {
	r io.Reader
}

func(r13 *rot13Reader) Read (b []byte) (int, error){
	n, err := r13.r.Read(b)
	for i:=range make([]byte, n){
		b[i] = rot13substitution(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	d := byte('L')

	fmt.Println(string(d))
	fmt.Println(string(rot13substitution(d)))
	fmt.Println(string(rot13substitution(rot13substitution(d))))

}
