package main

import (
	"fmt"
	"github.com/levinion/grr"
	"os"
)

// version: 0.1.0
func main() {
	Readfile("hello.md").Expect(func(v []byte) {
		fmt.Println("print result:", string(v))
	}).Else(func(err error) {
		fmt.Println("print err:", err)
	})
}

func Readfile(file string) *grr.Result[[]byte] {
	return grr.Try[[]byte](func(h *grr.Handler[[]byte]) {
		file, err := os.Open(file)
		h.Err(err)
		buffer := make([]byte, 4096)
		n, err := file.Read(buffer)
		h.Err(err)
		content := buffer[:n]
		h.OK(content)
	})
}
