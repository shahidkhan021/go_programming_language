package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type WriterLog struct{}

func main() {
	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	wl := WriterLog{}
	io.Copy(wl, resp.Body)
}

func (WriterLog) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Total Number of bytes return:", len(bs))
	return len(bs), nil
}
