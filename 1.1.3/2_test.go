package e1_1_3

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("1425"))

	average, error := CalculateAverage(reader)

	assert.Nil(t, error)
	assert.Equal(t, 3.0, average)
}

func TestCalculateAverageOneNumber(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("1"))

	average, error := CalculateAverage(reader)

	assert.Nil(t, error)
	assert.Equal(t, 1.0, average)
}

func TestCalculateAverageSameNumber(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("5555555555"))

	average, error := CalculateAverage(reader)

	assert.Nil(t, error)
	assert.Equal(t, 5.0, average)
}
