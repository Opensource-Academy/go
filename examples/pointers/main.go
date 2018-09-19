/*

In this small script we will give you some insight on how pointers work.
Dave Cheney has a great blogpost on what pointers are and how they work:
https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back

The code example below will hopefully give you some insight on how pointers work.
The use of struct and functions are intentional because this is way pointers are most offen used.

*/

package main

import (
	"fmt"
)

// New type Person
type Person struct {
	Name string
	Age  int
}

// function sets age to 35. Or does it...
func age41(p Person) {
	p.Age = 41
	fmt.Println("Person's age set to", p.Age)
}

// same function bu now with pointer
func age32(p *Person) {
	p.Age = 32
	fmt.Println("Person's age set to", p.Age)
}

func main() {
	// create person ivo
	ivo := Person{
		Name: "Ivo",
		Age:  40,
	}

	// let's print that
	fmt.Println("Name: ", ivo.Name, "initial age: ", ivo.Age)

	// set age ivo by calling the age31 function.
	age41(ivo)

	// ivo.Age was set to 35. Check if that's right.
	fmt.Println("Name: ", ivo.Name, "new age: ", ivo.Age) // ivo.Age is still 30...
	fmt.Println("-------------------------")
	// new person niels
	niels := Person{
		Name: "Niels",
		Age:  31,
	}

	// let's check
	fmt.Println("Name: ", niels.Name, "initial age: ", niels.Age)

	// Again, let's try to change the age but now sending out variable using a pointer "&"

	age32(&niels)
	fmt.Println("Name: ", niels.Name, "new age: ", niels.Age) // Succes!

	// The pointer "&" points to the memory location where niels is stored.
	// The asterisk "*" refers to the value that points to the given memory location.
}
