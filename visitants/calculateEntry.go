package visitants

import (
	"fmt"

	"github.com/cassiusbessa/gozoofunctions/zoostruct"
)

var zooData = zoostruct.GetZooStruct()

func CalculateEntry(person CountEntrantOld) float64 {
	var money float64
	var prices = zooData.Price
	fmt.Printf("%+v\n", zooData)
	fmt.Println(prices)
	fmt.Println(person)
	money += float64(person.child)*prices.Child + float64(person.adult)*prices.Adult + float64(person.senior)*prices.Senior
	return money

}
