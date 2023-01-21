package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reverse_battleships/battleships"
)

func main() {
	game := battleships.NewGame()
	ebiten.SetWindowSize(battleships.ScreenWidth, battleships.ScreenHeight)
	ebiten.SetWindowTitle("Reverse Battleships")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
