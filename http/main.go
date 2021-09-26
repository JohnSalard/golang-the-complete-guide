package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {

	}
	fmt.Println(res.Status)
	fmt.Println(res.StatusCode)

}
