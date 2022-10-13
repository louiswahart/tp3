package comsoc

func CopelandSWF(p Profile) (Count, error) {
	c := make(Count)

	err := checkProfileAlternative(p, p[0])

	if err == nil {
		for i, a := range p[0][:len(p[0])-1] {
			for _, b := range p[0][i+1:] {
				var count_a, count_b int
				for _, tab := range p {
					for _, alt := range tab {
						if alt == a {
							count_a++
							break
						} else if alt == b {
							count_b++
							break
						}
					}
				}
				if count_a > count_b {
					c[a]++
					c[b]--
				} else if count_b > count_a {
					c[b]++
					c[a]--
				}
			}
		}
	}
	return c, err
}

func CopelandSCF(p Profile) (bestAlts []Alternative, err error) {
	c, err := CopelandSWF(p)
	if err == nil {
		bestAlts = maxCount(c)
		return
	}
	return
}
