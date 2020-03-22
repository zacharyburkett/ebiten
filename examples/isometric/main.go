package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	risometric "github.com/hajimehoshi/ebiten/examples/resources/images/isometric"
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

	mapWidth  = 10
	mapHeight = 10
)

var (
	// viewport offsets
	camX, camY float64

	plr   player
	tiles [mapWidth][mapHeight]int

	tileImg               *ebiten.Image
	tileWidth, tileHeight int
)

func init() {
	plr.speed = 4

	// load images
	img, _, err := image.Decode(bytes.NewReader(rplatformer.MainChar_png))
	if err != nil {
		panic(err)
	}
	plr.img, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	img, _, err = image.Decode(bytes.NewReader(risometric.Tiles_png))
	if err != nil {
		panic(err)
	}
	tileImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h := tileImg.Size()
	tileWidth = w / 2
	tileHeight = h

	rand.Seed(time.Now().Unix())

	// randomize map tiles
	for i := 0; i < mapWidth; i++ {
		for j := 0; j < mapHeight; j++ {
			tiles[i][j] = rand.Intn(2)
		}
	}
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

	// draw tiles
	for i := 0; i < mapWidth*mapHeight; i++ {
		for j := 0; j <= i; j++ {
			if !(i-j < mapWidth && j < mapHeight) {
				continue
			}

			x := float64((j-i/2)*tileWidth - (tileWidth*(i%2))/2)
			y := float64(i * tileHeight / 2)

			var tile *ebiten.Image
			if tiles[i-j][j] == 0 {
				tile = tileImg.SubImage(image.Rect(0, 0, tileWidth, tileHeight)).(*ebiten.Image)
			} else {
				tile = tileImg.SubImage(image.Rect(tileWidth, 0, tileWidth*2, tileHeight)).(*ebiten.Image)
			}

			var op ebiten.DrawImageOptions
			op.GeoM.Translate(x-camX, y-camY)
			screen.DrawImage(tile, &op)
		}
	}

	// draw player
	var op ebiten.DrawImageOptions
	w, h := plr.img.Size()
	op.GeoM.Translate(plr.x-float64(w/2)-camX, plr.y-float64(h/2)-camY)
	screen.DrawImage(plr.img, &op)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Isometric Example"); err != nil {
		log.Fatal(err)
	}
}
