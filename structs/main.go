package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	//fmt.Println(resp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b := make([]byte, 999999)
	n, _ := resp.Body.Read(b)

	fmt.Println(string(b))
	fmt.Print(n)

}
