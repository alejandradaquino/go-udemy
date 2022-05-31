package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	fmt.Println(resp)
	if err != nil {
		fmt.Errorf("Error", err)
		os.Exit(1)
	}
}
