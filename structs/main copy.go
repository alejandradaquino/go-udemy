package main

import "fmt"

func mainFirstIteration() {
	alex := employee{"Alex", "Anderson", contactInfo{"email", 12}}
	println(alex.firstName)

	ale := owner{
		firstName: "Ale",
		lastName:  "Juarez",
		contactInfo: contactInfo{"email",
			12,
		},
	}
	ale.setFirstName("Alejandra")
	//ale.print()

	var juan employee
	juan.setFirstName("bla")
	//default empty strings
	//println(juan.firstName + " " + juan.lastName)

	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#00ff00",
		"white": "#ffffff",
		"green": "#0000ff",
	}

	for c, hex := range colors {
		fmt.Printf("value c %+v hex %+v\n", c, hex)
	}

	color2 := make(map[string]string)

	color2["white"] = "blabla"

	delete(colors, "red")

	fmt.Println(colors)
	fmt.Println(color2)

}
