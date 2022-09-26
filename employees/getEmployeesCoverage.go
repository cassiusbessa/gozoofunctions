package employees

import (
	"github.com/cassiusbessa/gozoofunctions/species"
	"github.com/cassiusbessa/gozoofunctions/zoostruct"
)

type Coverage struct {
	id, fullName string
	species      []string
	locations    []string
}

type SearchCoverage struct {
	Name, Id string
}

func GetOneCoverage(search SearchCoverage) Coverage {
	var employ zoostruct.Employe
	var coverage Coverage
	var speciesCoverage []zoostruct.Specie
	if search.Id != "" {
		employ = GetById(search.Id)
	} else {
		employ = GetByName(search.Name)
	}
	speciesCoverage = species.Byid(employ.ResponsibleFor)
	coverage.id = employ.Id
	coverage.fullName = employ.FirstName + " " + employ.LastName

	for _, s := range speciesCoverage {
		coverage.species = append(coverage.species, s.Name)
		coverage.locations = append(coverage.locations, s.Location)
	}

	return coverage
}
