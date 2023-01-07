package e1_1_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBiggestFirst(t *testing.T) {
	assert.Equal(t, 3, Biggest(3, 2, 1))
}

func TestBiggestSecond(t *testing.T) {
	assert.Equal(t, 3, Biggest(1, 3, 1))
}

func TestBiggestThird(t *testing.T) {
	assert.Equal(t, 3, Biggest(1, 2, 3))
}
