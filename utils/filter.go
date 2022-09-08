package utils

import (
	"github.com/cassiusbessa/gozoofunctions/zoostruct"
)

var zooData = zoostruct.GetZooStruct()

type AnimalFilter struct{ Specie, Sex string }

func FilterBySpecieAndSex(filter AnimalFilter) zoostruct.Specie {
	var specieFiltered zoostruct.Specie
	p := &specieFiltered
	for _, specie := range zooData.Species {
		if specie.Name == filter.Specie {
			*p = specie
		}
	}

	if filter.Sex != "" {
		var sexResidents []zoostruct.Resident
		for _, resident := range specieFiltered.Residents {
			if resident.Sex == filter.Sex {
				sexResidents = append(sexResidents, resident)
			}
		}
		specieFiltered.Residents = sexResidents
	}
	return specieFiltered

}
