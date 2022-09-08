package main

import (
	"fmt"

	// species "github.com/cassiusbessa/gozoofunctions/species"
	// employees "github.com/cassiusbessa/gozoofunctions/employees"
	// utils "github.com/cassiusbessa/gozoofunctions/utils"
	// animals "github.com/cassiusbessa/gozoofunctions/animals"

	"github.com/cassiusbessa/gozoofunctions/visitants"
)

func main() {
	// t := species.Byid([]string{"0938aa23-f153-4937-9f88-4858b24d6bce", "e8481c1d-42ea-4610-8e11-1752cfc05a46"})
	// t := species.GetAnimalsOlderThan("lions", 12)
	p := []visitants.Entrant{
		{Name: "Lara Carvalho", Age: 5},
		{Name: "Frederico Moreira", Age: 5},
		{Name: "Pedro Henrique Carvalho", Age: 5},
		{Name: "Maria Costa", Age: 18},
		{Name: "Núbia Souza", Age: 18},
		{Name: "Carlos Nogueira", Age: 50},
	}
	t := visitants.Count(p)
	a := visitants.CalculateEntry(t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%+v\n", a)

}
