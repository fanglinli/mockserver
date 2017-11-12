package mocker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandom(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Empty(returnRandom("there %is no %accepted for%mat str%ing % in this:"))

	randomValues := returnRandom("%s %d %d %s")
	assert.Equal(4, len(randomValues))
	_, ok := randomValues[0].(string)
	assert.True(ok)
	_, ok = randomValues[1].(int)
	assert.True(ok)
	_, ok = randomValues[2].(int)
	assert.True(ok)
	_, ok = randomValues[3].(string)
	assert.True(ok)
}

func TestRandomString(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Empty(randomString(-1))
	assert.Empty(randomString(0))
	assert.Equal(1, len(randomString(1)))
	assert.Equal(5, len(randomString(5)))
	assert.Equal(1337, len(randomString(1337)))
}

func TestGetSize(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	size, i := getSize(1, "%d")
	assert.Equal(5, size)
	assert.Equal(1, i)

	size, i = getSize(1, "%s")
	assert.Equal(5, size)
	assert.Equal(1, i)

	size, i = getSize(1, "%10s")
	assert.Equal(10, size)
	assert.Equal(3, i)

	size, i = getSize(1, "%8d")
	assert.Equal(8, size)
	assert.Equal(2, i)
}
