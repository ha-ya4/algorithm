package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSequentialSerch(t *testing.T) {
	s := []int{4, 5, 2, 10, 45, 1, 100}

	idx := intSequentialSerch(s, 10)
	assert.True(t, idx == 3)

	idx = intSequentialSerch(s, 10000)
	assert.True(t, idx == -1)
}