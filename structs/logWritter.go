package main

import "fmt"

type logWritter struct{}

func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
