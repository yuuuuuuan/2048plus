package main

import (
	game2048 "2048/2048"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	game, err := game2048.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(game2048.ScreenWidth, game2048.ScreenHeight)
	ebiten.SetWindowTitle("2048")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
