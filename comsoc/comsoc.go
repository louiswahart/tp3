package comsoc

import (
	"errors"
	"sort"
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
	keys := make([]Alternative, 0, len(count))
	for a := range count {
		keys = append(keys, a)
	}

	sort.SliceStable(keys, func(i, j int) bool { return count[keys[i]] > count[keys[j]] })

	max_c := count[keys[0]]
	for i := range keys {
		if count[keys[i]] != max_c {
			return keys[:i]
		}
	}
	return keys
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
