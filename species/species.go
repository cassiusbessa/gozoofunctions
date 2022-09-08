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

func GetAnimalsOlderThan(name string, old int) bool {
	var animal []zoostruct.Resident
	var older = true
	p := &older
	for _, specie := range zooData.Species {
		if specie.Name == name {
			animal = append(animal, specie.Residents...)
		}
	}
	for _, resident := range animal {
		if resident.Age < old {
			*p = false
		}
	}
	return older
}
