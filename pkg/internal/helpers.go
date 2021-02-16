package internal

import "strconv"

func getRuneNumber(id string) int {
	runeNumber, err := strconv.ParseInt(id[1:], 10, 0)
	if err != nil {
		return 0
	}
	return int(runeNumber)
}
