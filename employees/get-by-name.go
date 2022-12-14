package employees

import (
	"github.com/cassiusbessa/gozoofunctions/zoostruct"
)

var zooData = zoostruct.GetZooStruct()

func GetByName(name string) zoostruct.Employe {
	var employ zoostruct.Employe
	for _, e := range zooData.Employees {
		if e.FirstName == name || e.LastName == name {
			employ = e
		}
	}
	return employ
}
