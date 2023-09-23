package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(resp)

	// bs := make([]byte, 99999)
	// n, err := resp.Body.Read(bs)
	// fmt.Println(n)
	// fmt.Println(string(bs))

	lw := logWriter{}

	// the above code is equivalent to the below line
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Print((string(bs)))
	fmt.Print("Just wrote this many bytes: ")
	return len(bs), nil
}
