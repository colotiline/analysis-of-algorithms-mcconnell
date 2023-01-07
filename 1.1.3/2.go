package e1_1_3

import (
	"bufio"
	"io"
)

func CalculateAverage(reader *bufio.Reader) (average float64, error error) {
	count, average := 0.0, 0

	for {
		if rune, _, error := reader.ReadRune(); error != nil {
			if error == io.EOF {
				return average, nil
			}

			return 0, error
		} else {
			count++

			average = average + (float64(rune-'0')-average)/count
		}
	}
}
