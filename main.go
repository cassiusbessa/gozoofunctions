package main

import (

	// species "github.com/cassiusbessa/gozoofunctions/species"
	// employees "github.com/cassiusbessa/gozoofunctions/employees"
	// utils "github.com/cassiusbessa/gozoofunctions/utils"

	"fmt"

	"github.com/cassiusbessa/gozoofunctions/employees"
)

func main() {
	// t := species.Byid([]string{"0938aa23-f153-4937-9f88-4858b24d6bce", "e8481c1d-42ea-4610-8e11-1752cfc05a46"})
	// animals.WeekSchedule()
	// o := animals.Options{IncludeNames: true, Sorted: true}
	// t := animals.GetMap(o)
	// for k, v := range animals.GetSchedule() {
	// 	fmt.Println(k, v)
	// }
	f := employees.SearchCoverage{Id: "fdb2543b-5662-46a7-badc-93d960fdc0a8"}
	a := employees.GetOneCoverage(f)
	fmt.Printf("%+v\n", a)

}
