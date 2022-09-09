package animals

type Options struct {
	IncludeNames bool
	Sex          string
	Sorted       bool
}

func GetMap(options Options) interface{} {
	animalsMap := make(map[string]map[string][]string)
	cardenal := []string{"NE", "NW", "SE", "SW"}
	for _, c := range cardenal {
		animalsMap[c] = map[string][]string{}
	}
	for _, s := range zooData.Species {
		for _, r := range s.Residents {
			animalsMap[s.Location][s.Name] = append(animalsMap[s.Location][s.Name], r.Name)
		}
	}
	return animalsMap
}
