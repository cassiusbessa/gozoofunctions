package main

import (
	"fmt"

	// species "github.com/cassiusbessa/gozoofunctions/species"
	// employees "github.com/cassiusbessa/gozoofunctions/employees"
	// utils "github.com/cassiusbessa/gozoofunctions/utils"
	animals "github.com/cassiusbessa/gozoofunctions/animals"
)

func main() {
	// t := species.Byid([]string{"0938aa23-f153-4937-9f88-4858b24d6bce", "e8481c1d-42ea-4610-8e11-1752cfc05a46"})
	// t := species.GetAnimalsOlderThan("lions", 12)
	o := animals.Options{IncludeNames: true, Sex: "female"}
	t := animals.GetMap(o)
	fmt.Printf("%+v\n", t)
	// fmt.Printf("%+v\n", a)

}
