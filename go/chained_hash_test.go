package algorithm

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashValue(t *testing.T) {
	c := newChainedHash(10)

	key := 5
	hash, err := c.hashValue(key)
	assert.NoError(t, err)

	for i := 0; i < 10; i++ {
		h, err := c.hashValue(key)
		assert.NoError(t, err)
		assert.True(t, hash == h)
	}

	key2 := "hello"
	hash, err = c.hashValue(key2)
	assert.NoError(t, err)

	for i := 0; i < 10; i++ {
		h, err := c.hashValue(key2)
		assert.NoError(t, err)
		assert.True(t, hash == h)
	}

	key3 := true
	_, err = c.hashValue(key3)
	assert.Exactly(t, err, errors.New(errUnknownType))
}

func TestAdd(t *testing.T) {
	c := newChainedHash(10)

	target := 66
	targeth, err := c.hashValue(target)
	tvalue := 1000
	assert.NoError(t, err)
	err = c.add(target, tvalue)
	assert.NoError(t, err)

	key := 6
	value := 100
	assert.NoError(t, err)
	err = c.add(key, value)
	assert.NoError(t, err)

	assert.True(t, c.table[targeth].next.value == tvalue)
	assert.Exactly(t, c.add(target, tvalue), errors.New(errKeyExists))
}

func TestSerch(t *testing.T) {
	c := newChainedHash(10)

	key := 6
	value := 100
	err := c.add(key, value)
	assert.NoError(t, err)

	v, err := c.serch(key)
	assert.NoError(t, err)
	assert.True(t, v == value)

	_, err = c.serch("hello")
	assert.Exactly(t, err, errors.New(errNoExist))
}

func TestRemove(t *testing.T) {
	c := newChainedHash(10)

	target := 66
	tvalue := 1000
	err := c.add(target, tvalue)
	assert.NoError(t, err)

	v, err := c.serch(target)
	assert.NoError(t, err)
	assert.True(t, v == tvalue)

	err = c.remove(target)
	assert.NoError(t, err)
	_, err = c.serch(target)
	assert.Exactly(t, err, errors.New(errNoExist))

	err = c.remove(target)
	assert.Exactly(t, err, errors.New(errKeyNotExist))
}
