package e1_1_3

import (
	"fmt"
	"sort"
)

func SecondSmallest(values []int) (value int, error error) {
	if len(values) < 2 {
		return 0, fmt.Errorf(
			"minimum length is 2, your is %v",
			len(values),
		)
	}

	sort.Ints(values)

	return values[1], nil
}
