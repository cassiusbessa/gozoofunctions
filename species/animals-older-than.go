package species

import (
	data "github.com/cassiusbessa/gozoofunctions/zoostruct"
)

func GetAnimalsOlderThan(name string, old int) bool {
	var animal []data.Resident
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
