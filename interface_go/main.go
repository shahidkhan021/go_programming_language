package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// net/http will call a interface and inject response of http
	// to a byte slice through a reader function we need to initialize
	// a byte slice of bigger size

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	/* condensed code of the above */
	io.Copy(os.Stdout, resp.Body)

}
