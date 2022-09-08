package zoostruct

import (
	"encoding/json"
	"os"
)

type zoo_data struct {
	Species   []Specie  `json:"species"`
	Employees []employe `json:"employees"`
	Hours     hour      `json:"hours"`
	Price     price     `json:"price"`
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

type employe struct {
	Id             string   `json:"id"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Managers       []string `json:"managers"`
	ResponsibleFor []string `json:"responsibleFor"`
}

type hour struct {
	Tuesday struct {
		Open  string `json:"open"`
		Close string `json:"close"`
	}
	Wednesday struct {
		Open  string `json:"open"`
		Close string `json:"close"`
	}
	Thursday struct {
		Open  string `json:"open"`
		Close string `json:"close"`
	}
	Friday struct {
		Open  string `json:"open"`
		Close string `json:"close"`
	}
	Saturday struct {
		Open  string `json:"open"`
		Close string `json:"close"`
	}
	Sunday struct {
		Open  string `json:"open"`
		Close string `json:"close"`
	}
	Monday struct {
		Open  string `json:"open"`
		Close string `json:"close"`
	}
}

type price struct {
	Adult  float32 `json:"adult"`
	Child  float32 `json:"child"`
	Senior float32 `json:"senior"`
}

func GetZooStruct() zoo_data {
	data, error := os.ReadFile("./data/zoo_data.json")
	if error != nil {
		panic(error)
	}
	zoo := zoo_data{}
	json.Unmarshal([]byte(data), &zoo)
	return zoo

}
