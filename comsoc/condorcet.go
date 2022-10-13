package comsoc

func CondorcetWinner(p Profile) (bestAlts []Alternative, err error) {
	err = checkProfileAlternative(p, p[0])
	if err == nil {
		bestAlts = make([]Alternative, 0, len(p[0]))
		for _, candidate := range p[0] {
			var stillCandidate bool = true
			for _, opponent := range p[0] {
				if candidate != opponent {
					var candidateCount, opponentCount int
					for _, prefs := range p {
						for _, alt := range prefs {
							if alt == candidate {
								candidateCount += 1
								break
							} else if alt == opponent {
								opponentCount += 1
								break
							}
						}
					}
					if opponentCount >= candidateCount {
						stillCandidate = false
						break
					}
				}
			}
			if stillCandidate {
				bestAlts = append(bestAlts, candidate)
			}
		}
	}
	return
}
