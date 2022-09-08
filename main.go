package main

import (
	"fmt"

	species "github.com/cassiusbessa/gozoofunctions/species"
)

func main() {
	// t := species.Byid([]string{"0938aa23-f153-4937-9f88-4858b24d6bce", "e8481c1d-42ea-4610-8e11-1752cfc05a46"})
	t := species.GetAnimalsOlderThan("lions", 7)
	fmt.Printf("%+v\n", t)
}
