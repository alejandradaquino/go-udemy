package main

import "fmt"

func main() {
	alex := person{"Alex", "Anderson", contactInfo{"email", 12}}
	println(alex.firstName)

	ale := person{
		firstName: "Ale",
		lastName:  "Juarez",
		contactInfo: contactInfo{"email",
			12,
		},
	}
	ale.setFirstName("Alejandra")
	ale.print()

	var juan person
	//default empty strings
	println(juan.firstName + " " + juan.lastName)

	colors := map[string]string{
		"red": "#fff0000",
	}
	fmt.Println(colors)

}
