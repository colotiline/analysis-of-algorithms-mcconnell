package e1_1_3

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestZeroRunes(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(""))

	number, error := CountCapitalLetters(reader)

	assert.Zero(t, number)
	assert.Nil(t, error)
}

func TestOneCapitalLetter(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("12bdjE!"))

	number, error := CountCapitalLetters(reader)

	assert.Equal(t, 1, number)
	assert.Nil(t, error)
}

func TestAllCapitalLetters(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("BHGFOKJDH"))

	number, error := CountCapitalLetters(reader)

	assert.Equal(t, 9, number)
	assert.Nil(t, error)
}

func TestCyrillicCapitalLetters(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("BHGFOKJDHАБч1213!"))

	number, error := CountCapitalLetters(reader)

	assert.Equal(t, 11, number)
	assert.Nil(t, error)
}
