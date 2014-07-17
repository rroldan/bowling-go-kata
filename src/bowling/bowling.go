package bowling

type Game struct {
score int

}
func (game *Game) Roll(pins int) {
    
}

func (game *Game) Score() int {
	return game.score
}

