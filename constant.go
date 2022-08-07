package main

import "fmt"

func main() {
	const firstName = "John"
	const lastName = "Doe"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		age     int8 = 19
		address      = "Tulungagung"
	)

	fmt.Println(age)
	fmt.Println(address)
}
