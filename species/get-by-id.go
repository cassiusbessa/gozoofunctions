package species

import (
	"github.com/cassiusbessa/gozoofunctions/zoostruct"
)

var zooData = zoostruct.GetZooStruct()

func Byid(id []string) []zoostruct.Specie {
	var animal []zoostruct.Specie
	for _, id := range id {
		for _, specie := range zooData.Species {
			if specie.Id == id {
				animal = append(animal, specie)
			}
		}
	}
	return animal
}
