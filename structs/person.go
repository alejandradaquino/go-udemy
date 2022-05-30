package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {

	fmt.Printf("%+v", p)
}

func (p *person) setFirstName(name string) {
	(*p).firstName = name
}
