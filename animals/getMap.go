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

	if options.IncludeNames {
		for _, s := range zooData.Species {
			for _, r := range s.Residents {
				if options.Sex == "" {
					animalsMap[s.Location][s.Name] = append(animalsMap[s.Location][s.Name], r.Name)
				} else {
					if r.Sex == options.Sex {
						animalsMap[s.Location][s.Name] = append(animalsMap[s.Location][s.Name], r.Name)
					}
				}
			}
		}
	}
	return animalsMap
}
