package comsoc

func MajoritySWF(p Profile) (count Count, err error) {
	count = make(Count)
	err = checkProfileAlternative(p, p[0])
	if err == nil {
		for _, prof := range p {
			count[prof[0]] += 1
		}
	}
	return
}

func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := MajoritySWF(p)
	if err == nil {
		bestAlts = maxCount(count)
	}
	return
}
