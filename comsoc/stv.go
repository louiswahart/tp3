package comsoc

func STV_SWF(p Profile) (count Count, err error) {
	count = make(Count)
	err = checkProfileAlternative(p, p[0])

	if err == nil {
		var outCandidates []Alternative
		for i := 0; i < len(p[0])-1; i++ {
			var low, high Alternative
			var c Count = make(Count)
			for _, prefs := range p {
				for _, alt := range prefs {
					if !contains(outCandidates, alt) {
						c[alt]++
						low = alt
						high = alt
						break
					}
				}
			}
			for k, v := range c {
				if v > c[high] {
					high = k
				} else if v < c[low] {
					low = k
				}
			}
			if c[high] >= len(c)/2 {
				count[high] = i + 2
				for _, alt := range p[0] {
					if count[alt] == 0 {
						count[alt] = i + 1
					}
				}
				break
			} else {
				count[low] = i + 1
				outCandidates = append(outCandidates, low)
			}
		}
	}
	return count, err
}

func STV_SCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := STV_SWF(p)

	if err == nil {
		bestAlts = maxCount(count)
	}
	return bestAlts, err
}
