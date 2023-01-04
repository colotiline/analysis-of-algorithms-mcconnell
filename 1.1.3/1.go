package e1_1_3

import (
	"bufio"
	"io"
	"unicode"
)

func CountCapitalLetters(reader *bufio.Reader) (number int, error error) {
	number = 0

	for {
		if rune, _, error := reader.ReadRune(); error != nil {
			if error == io.EOF {
				return number, nil
			}

			return 0, error
		} else {
			if unicode.IsUpper(rune) {
				number++
			}
		}
	}
}
