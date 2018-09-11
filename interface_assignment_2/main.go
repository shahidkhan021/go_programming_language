package main

import (
	"fmt"
	"io"
	"os"
)

type writerLog struct{}

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	wl := writerLog{}
	io.Copy(wl, f)
}

func (writerLog) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	return len(b), nil
}
