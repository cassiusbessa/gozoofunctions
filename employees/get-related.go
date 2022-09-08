package employees

import "errors"

func GetRelated(id string) ([]string, error) {
	var employ []string
	for _, e := range zooData.Employees {
		for _, m := range e.Managers {
			if m == id {
				employ = append(employ, e.FirstName+" "+e.LastName)
			}
		}
	}
	if len(employ) == 0 {
		return employ, errors.New("o id inserido não é de uma pessoa colaboradora gerente")
	}
	return employ, nil
}
