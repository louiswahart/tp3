package comsoc

import (
	"errors"
)

type Alternative int
type Profile [][]Alternative
type Count map[Alternative]int

func rank(alt Alternative, prefs []Alternative) int {
	for i, v := range prefs {
		if v == alt {
			return i
		}
	}
	return -1
}

func isPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	for _, v := range prefs {
		if v == alt1 {
			return true
		}
		if v == alt2 {
			return false
		}
	}
	return false
}

func maxCount(count Count) (bestAlts []Alternative) {
	max := -1
	for k := range count {
		if count[k] > max {
			max = count[k]
			bestAlts = make([]Alternative, 1)
			bestAlts[0] = k
		} else if count[k] == max {
			bestAlts = append(bestAlts, k)
		}
	}
	return bestAlts
}

func checkProfileAlternative(prefs Profile, alts []Alternative) error {
	for _, profile := range prefs {
		for _, val := range profile {
			appeared := false
			if len(profile) != len(alts) {
				return errors.New("Lengths don't match")
			}
			for _, alt := range alts {
				if val == alt {
					if appeared == true {
						return errors.New("Duplicate Alternative")
					} else {
						appeared = true
					}
				}
			}
			if appeared == false {
				return errors.New("Missing Alternative")
			}
		}
	}
	return nil
}

func contains(alts []Alternative, alt Alternative) bool {
	for _, a := range alts {
		if a == alt {
			return true
		}
	}
	return false
}
