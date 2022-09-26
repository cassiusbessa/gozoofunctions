package species

import (
	"fmt"

	"github.com/cassiusbessa/gozoofunctions/zoostruct"
)

type Output struct {
	name, sex string
	age       int
}

func Oldest(specie zoostruct.Specie) Output {
	fmt.Println(specie)
	var oldest = Output{}
	var actualAge = 0
	for _, s := range specie.Residents {
		fmt.Println(actualAge, s.Age)
		if actualAge < s.Age {
			actualAge = s.Age
			oldest.name = s.Name
			oldest.sex = s.Sex
			oldest.age = s.Age
		}
	}
	return oldest
}
