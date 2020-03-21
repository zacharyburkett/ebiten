package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	rplatformer "github.com/hajimehoshi/ebiten/examples/resources/images/platformer"
)

type player struct {
	x, y  float64
	speed float64

	img *ebiten.Image
}

var (
	// viewport offsets
	camX, camY float64

	plr player
)

func init() {
	plr.speed = 2

	img, _, err := image.Decode(bytes.NewReader(rplatformer.MainChar_png))
	if err != nil {
		panic(err)
	}
	plr.img, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func update(screen *ebiten.Image) error {
	// player movement inputs
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		plr.y -= plr.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		plr.y += plr.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		plr.x -= plr.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		plr.x += plr.speed
	}

	// skip frame if running slowly
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(plr.x, plr.y)
	screen.DrawImage(plr.img, op)

	return nil
}

func main() {
	if err := ebiten.Run(update, 1024, 512, 1, "Isometric Example"); err != nil {
		log.Fatal(err)
	}
}
