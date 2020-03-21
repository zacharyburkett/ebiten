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

const (
	screenWidth  = 1024
	screenHeight = 512
)

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

	// update camera offset to follow player
	camX, camY = plr.x-(screenWidth/2), plr.y-(screenHeight/2)

	// skip frame if running slowly
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(plr.x-camX, plr.y-camY)
	screen.DrawImage(plr.img, op)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Isometric Example"); err != nil {
		log.Fatal(err)
	}
}
