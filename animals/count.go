package animals

import (
	utils "github.com/cassiusbessa/gozoofunctions/utils"
	"github.com/cassiusbessa/gozoofunctions/zoostruct"
)

type counted struct {
	species string
	total   int
}

var zooData = zoostruct.GetZooStruct()

func CountResident(filter utils.AnimalFilter) counted {
	var count counted
	p := &count
	if filter.Specie == "" {
		for _, s := range zooData.Species {
			p.species = "All Species"
			p.total += len(s.Residents)
		}
	} else {
		filtered := utils.FilterBySpecieAndSex(filter)
		p.species = filtered.Name
		p.total += len(filtered.Residents)
	}
	return count
}
