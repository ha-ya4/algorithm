package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntBinarySerch(t *testing.T) {
	s := []int{1, 2, 5, 10, 57, 60, 76, 77, 83, 90, 99}

	// 左側
	idx := intBinarySerch(s, 5)
	assert.True(t, idx == 2)

	// 右側
	idx = intBinarySerch(s, 77)
	assert.True(t, idx == 7)

	// なし
	idx = intBinarySerch(s, 1000)
	assert.True(t, idx == -1)
}