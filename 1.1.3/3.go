package e1_1_3

func AllDifferent(first int, second int, third int) bool {
	if first != second && first != third && second != third {
		return true
	}

	return false
}
