package employees

import (
	"github.com/cassiusbessa/gozoofunctions/zoostruct"
)

func GetById(id string) zoostruct.Employe {
	var employ zoostruct.Employe
	for _, e := range zooData.Employees {
		if e.Id == id {
			employ = e
		}
	}
	return employ

}
