package comsoc

func ApprovalSWF(p Profile, thresholds []int) (count Count, err error) {
	count = make(Count)
	err = checkProfileAlternative(p, p[0])
	if err == nil {
		for i, prof := range p {
			for j, val := range prof {
				if j < thresholds[i] {
					count[val] += 1
				}
			}
		}
	}
	return
}
func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	count, err := BordaSWF(p)
	if err == nil {
		bestAlts = maxCount(count)
	}
	return
}
