package main

import (
	"fmt"
	"ia04/comsoc"
)

func main() {
	size := 3
	prefs := make(comsoc.Profile, size)
	for i := 0; i < size; i++ {
		prefs[i] = make([]comsoc.Alternative, size)
	}
	prefs[0][0] = 1
	prefs[0][1] = 3
	prefs[0][2] = 2
	prefs[1][0] = 2
	prefs[1][1] = 3
	prefs[1][2] = 1
	prefs[2][0] = 1
	prefs[2][1] = 2
	prefs[2][2] = 3

	thresholds := make([]int, size)
	thresholds[0] = 3
	thresholds[1] = 2
	thresholds[2] = 1

	fmt.Println(comsoc.ApprovalSWF(prefs, thresholds))
	fmt.Println(comsoc.ApprovalSCF(prefs, thresholds))
}
