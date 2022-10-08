package comsoc

func BordaSWF(p Profile) (count Count, err error) {
	count = make(Count)
	err = checkProfileAlternative(p, p[0])
	if err == nil {
		for _, prof := range p {
			for i, val := range prof {
				count[val] += len(prof) - 1 - i
			}
		}
	}
	return
}

func BordaSCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := BordaSWF(p)
	if err == nil {
		bestAlts = maxCount(count)
	}
	return
}
