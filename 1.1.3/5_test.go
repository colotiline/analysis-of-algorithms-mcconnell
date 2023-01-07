package e1_1_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecondSmallest(t *testing.T) {
	value, error := SecondSmallest([]int{3, 1, 2})

	assert.Nil(t, error)
	assert.Equal(t, 2, value)
}

func TestSecondSmallestLengthIsTooSmall(t *testing.T) {
	value, error := SecondSmallest([]int{3})

	assert.Equal(t, "minimum length is 2, your is 1", error.Error())
	assert.Equal(t, 0, value)
}
