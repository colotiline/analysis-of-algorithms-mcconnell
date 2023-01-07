package e1_1_3

func Biggest(first int, second int, third int) int {
	if first > second {
		if first > third {
			return first
		} else {
			return third
		}
	}

	if second > third {
		return second
	}

	return third
}
