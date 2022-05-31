package main

import "fmt"

type person interface {
	getFirstName() string
	getLastName() string
	getContactInfo() contactInfo
}

type employee struct {
	firstName string
	lastName  string
	contactInfo
}

type owner struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p employee) print() {

	fmt.Printf("%+v", p)
}

func (p *employee) setFirstName(name string) {
	(*p).firstName = name
}

func (p owner) print() {

	fmt.Printf("%+v", p)
}

func (p *owner) setFirstName(name string) {
	(*p).firstName = name
}
