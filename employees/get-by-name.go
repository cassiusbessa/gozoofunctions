package employees

import (
	"github.com/cassiusbessa/gozoofunctions/zoostruct"
)

var zooData = zoostruct.GetZooStruct()

func GetByName(name string) []zoostruct.Employe {
	var employees []zoostruct.Employe
	for _, employ := range zooData.Employees {
		if employ.FirstName == name || employ.LastName == name {
			employees = append(employees, employ)
		}
	}
	return employees
}
