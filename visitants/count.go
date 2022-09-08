package visitants

type Entrant struct {
	Name string
	Age  int
}

type CountEntrantOld struct {
	child, adult, senior int
}

func Count(entrants []Entrant) CountEntrantOld {
	var counted CountEntrantOld
	p := &counted
	for _, e := range entrants {
		switch {
		case e.Age >= 50:
			p.senior += 1
		case e.Age < 18:
			p.child += 1
		default:
			p.adult += 1
		}
	}
	return counted
}
