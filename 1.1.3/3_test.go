package e1_1_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllDifferent(t *testing.T) {
	assert.True(t, AllDifferent(0, 1, 2))
}

func TestAllDifferentOneAndSecondAreSame(t *testing.T) {
	assert.False(t, AllDifferent(0, 0, 2))
}

func TestAllDifferentOneAndThirdAreSame(t *testing.T) {
	assert.False(t, AllDifferent(0, 1, 0))
}

func TestAllDifferentSecondAndThirdAreSame(t *testing.T) {
	assert.False(t, AllDifferent(0, 1, 1))
}

func TestAllDifferentSecondAllSame(t *testing.T) {
	assert.False(t, AllDifferent(0, 0, 0))
}
