package main

import (
	_ "image/jpeg"
	"log"
	"math"

    "github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/ebitenutil"

)

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	count        = 0
	gophersImage *ebiten.Image
)

func update(screen *ebiten.Image) error {
	count++
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	w, h := gophersImage.Size()
	op := &ebiten.DrawImageOptions{}

	// Move the image's center to the screen's upper-left corner.
	// This is a preparation for rotating. When geometry matrices are applied,
	// the origin point is the upper-left corner.
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

	// Rotate the image. As a result, the anchor point of this rotate is
	// the center of the image.
	op.GeoM.Rotate(float64(count%360) * 2 * math.Pi / 360)

	// Move the image to the screen's center.
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	screen.DrawImage(gophersImage, op)
	return nil
}

func main() {
	gophersImage, _, _ = ebitenutil.NewImageFromFile("gfx/dick.png", ebiten.FilterDefault)

	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Rotate (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}