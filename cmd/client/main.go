package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hmdyt/orange/adapter/client"
	"github.com/hmdyt/orange/adapter/client/api"
)

func main() {
	apiClient := api.NewClient("")
	game := client.NewGame(apiClient)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("虫バトル")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
