package bowling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScore(t *testing.T) {

	game := Game{0}

	assert.Equal(t,  game.Score(), 0)
}

func TestRoll(t *testing.T){
	game := Game{0}
	for i:=0; i<20; i++ {
		game.Roll(0);
    }
	assert.Equal(t, game.Score(),0)

}
