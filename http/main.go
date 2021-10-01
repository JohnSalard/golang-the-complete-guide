package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)

	// fmt.Println(res.Body.Read(bs))

	// log.Info()
	lw := logWriter{}
	// io.Copy(os.Stdout, res.Body)
	io.Copy(lw, res.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
