package main

import "fmt"

func main() {
	var name string

	name = "John"
	fmt.Println(name)

	name = "John Doe"
	fmt.Println(name)

	var friendName = "Syekh havis"
	fmt.Println(friendName)

	var age int8 = 20
	fmt.Println(age)

	address := "Tulungagung"
	fmt.Println(address)

	var (
		firstName      = "John "
		lastName       = "Doe"
		noHp      int8 = 1424342
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(noHp)
}
