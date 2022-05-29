package main

func main() {
	alex := person{"Alex", "Anderson"}
	println(alex.firstName)

	ale := person{firstName: "Ale", lastName: "Juarez"}
	println(ale.firstName + " " + ale.lastName)
}
