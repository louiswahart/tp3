package comsoc

import "errors"

func TieBreakFactory(alts_order []Alternative) func([]Alternative) (Alternative, error) {
	return func(alts []Alternative) (best_alt Alternative, err error) {
		for _, pref := range alts_order {
			for _, alt := range alts {
				if pref == alt {
					return pref, err
				}
			}
		}
		err = errors.New("Did not find corresponding match in given Alts")
		return
	}
}
