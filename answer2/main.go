package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	green := color.RGBA{0, 255, 0, 255}
	red := color.RGBA{200, 30, 30, 255}
	minX, minY, maxX, maxY := 0, 0, 200, 200
	rectImg := image.NewRGBA(image.Rect(minX, minY, maxX, maxY))
	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP,
		draw.Src)

	interval := 10
	for x := interval; x < maxX; x += interval {
		for y := minY; y <= maxY; y++ {
			rectImg.Set(x, y, red)
		}
	}
	for y := interval; y < maxY; y += interval {
		for x := minX; x <= maxX; x++ {
			rectImg.Set(x, y, red)
		}
	}

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)

}
