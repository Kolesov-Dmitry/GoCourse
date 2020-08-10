package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

// fillRect draws filled rectangle with given color
// Input:
//  - img is the surface to draw the rectangle
//  - rect is the position and size of the rectanle
//  - pen is the color to fill the rectangle
func fillRect(img *image.RGBA, rect image.Rectangle, pen color.RGBA) {
	if img == nil {
		return
	}

	draw.Draw(img, rect, &image.Uniform{pen}, image.ZP, draw.Src)
}

// fillRect draws line from the start to the end point
// Input:
//  - img is the surface to draw the rectangle
//  - start is the start point of the line
//  - end is the end point of the line
//  - pen is the color to draw the line
func drawLine(img *image.RGBA, start, end image.Point, pen color.RGBA) {
	if img == nil || start == end {
		return
	}

	if start.X > end.X {
		start, end = end, start
	}

	if start.X != end.X {
		// Linear interpolation between the start and end points
		for x := start.X; x <= end.X; x++ {
			y := (start.Y*(end.X-x) + end.Y*(x-start.X)) / (end.X - start.X)

			img.Set(x, y, pen)
		}
	} else {
		// draw vertical line
		for y := start.Y; y <= end.Y; y++ {
			img.Set(start.X, y, pen)
		}
	}
}

// saveImageToFile saves img to the file
// Input:
//  - img is the image to save
//  - path is the name of the file
// Ouput:
//  - nil, in case of success
//  - otherwise, returns the occured error
func saveImageToFile(img *image.RGBA, path string) error {
	if img == nil || len(path) == 0 {
		return errors.New("saveToFile invalid arguments")
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

// Image size
const (
	imageWidth  = 200
	imageHeight = 200
)

// Colors
var teal = color.RGBA{0, 128, 128, 255}
var white = color.RGBA{255, 255, 255, 255}

func main() {
	// create image filled with teal color
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	fillRect(img, img.Bounds(), teal)

	// draw vertical lines
	var stepX = int(imageWidth / 5)
	for x := stepX; x < imageWidth; x += stepX {
		drawLine(img, image.Point{x, 0}, image.Point{x, imageHeight}, white)
	}

	// draw horizontal lines
	var stepY = int(imageHeight / 5)
	for y := stepY; y < imageHeight; y += stepY {
		drawLine(img, image.Point{0, y}, image.Point{imageWidth, y}, white)
	}

	// save image to file
	if err := saveImageToFile(img, "rectangle.png"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Image has been created successfully")
	}
}
