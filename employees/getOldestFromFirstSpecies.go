package employees

import (
	"github.com/cassiusbessa/gozoofunctions/species"
	"github.com/cassiusbessa/gozoofunctions/zoostruct"
)

func firstSpecie(employ zoostruct.Employe) string {
	return employ.ResponsibleFor[0]
}

func GetOldestFromFirstSpecies(id string) species.Output {
	var employ = GetById(id)
	var idFirstSpecie = firstSpecie(employ)
	var specie = species.Byid([]string{idFirstSpecie})
	return species.Oldest(specie[0])
}
