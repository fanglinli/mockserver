package mocker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseName(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("", baseName(""))
	assert.Equal("filename", baseName("filename"))
	assert.Equal("b", baseName("b.txt"))
	assert.Equal("automobiles", baseName("automobiles.json"))
	assert.Equal("location", baseName("location.current.csv"))
}
