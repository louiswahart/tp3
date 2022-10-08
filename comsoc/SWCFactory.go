package comsoc

func SWCFactory(swf func(p Profile) (Count, error), tiebreak func([]Alternative) (Alternative, error)) func(Profile) (Alternative, error) {
	return func(p Profile) (alts Alternative, err error) {
		c, err := swf(p)
		if err == nil {
			bestAlts := maxCount(c)
			bestAlt, err := tiebreak(bestAlts)
			return bestAlt, err
		}
		return
	}
}
