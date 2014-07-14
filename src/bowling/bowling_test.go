package bowling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScore(t *testing.T) {

	game := Game{0}

	assert.Equal(t,  game.Score(), 0)
}
