package main

import (
	"fmt"
	"os"
)

func main() {
	data, error := os.ReadFile("./data/zoo_data.json")
	if error != nil {
		panic(error)
	}
	fmt.Println(string(data))
}
