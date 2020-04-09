package main

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	//"image/color"
)

func main() {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 256, 256))
	gc := draw2dimg.NewGraphicContext(dest)

	// Draw a closed shape
	gc.BeginPath() // Initialize a new path
	gc.MoveTo(80, 30) // Move to a position to start the new path
	gc.LineTo(200, 90)
	gc.Close()
	gc.FillStroke()

	// Save to file
	draw2dimg.SaveToPngFile("hello.png", dest)
}
