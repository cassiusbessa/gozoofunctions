package zoostruct

import (
	"encoding/json"
	"os"
)

type zoo_data struct {
	Species   []Specie  `json:"species"`
	Employees []Employe `json:"employees"`
	Hours     Hour      `json:"hours"`
	Price     price     `json:"prices"`
}

type Specie struct {
	Id           string     `json:"id"`
	Name         string     `json:"name"`
	Popularity   int        `json:"popularity"`
	Location     string     `json:"location"`
	Availability []string   `json:"availability"`
	Residents    []Resident `json:"residents"`
}

type Resident struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
}

type Employe struct {
	Id             string   `json:"id"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Managers       []string `json:"managers"`
	ResponsibleFor []string `json:"responsibleFor"`
}

type Hour struct {
	Tuesday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	}
	Wednesday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	}
	Thursday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	}
	Friday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	}
	Saturday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	}
	Sunday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	}
	Monday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	}
}

type price struct {
	Adult  float64 `json:"adult"`
	Child  float64 `json:"child"`
	Senior float64 `json:"senior"`
}

func GetZooStruct() zoo_data {
	data, error := os.ReadFile("./data/zoo_data.json")
	if error != nil {
		panic(error)
	}
	zoo := zoo_data{}
	json.Unmarshal(data, &zoo)
	return zoo
}
