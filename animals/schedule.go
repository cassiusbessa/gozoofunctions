package animals

import (
	"encoding/json"
	"strconv"
)

type week struct {
	officeHour string
	exhibition []string
}

var weekDays = []string{"Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday", "Monday"}

func GetSchedule(day string) []string {
	animalExhbt := []string{}
	for _, s := range zooData.Species {
		for _, d := range s.Availability {
			if d == day {
				animalExhbt = append(animalExhbt, s.Name)
			}
		}
	}
	return animalExhbt
}

func GetHours(day string) string {
	h, err := json.Marshal(zooData.Hours)
	if err != nil {
		return err.Error()
	}
	hours := make(map[string]map[string]int)
	e := json.Unmarshal(h, &hours)
	if e != nil {
		return e.Error()
	}
	var dayHour string

	for k, v := range hours {
		if day == k {
			dayHour = "Open from " +
				strconv.Itoa(v["open"]) + "am until " + strconv.Itoa(v["close"]) + "pm"
		}
	}
	return dayHour
}

func WeekSchedule() map[string]week {
	schedule := make(map[string]week)
	for _, w := range weekDays {
		schedule[w] = week{GetHours(w), GetSchedule(w)}
	}
	return schedule
}
