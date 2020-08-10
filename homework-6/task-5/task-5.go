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

const (
	imageSize         = 600
	boardSise         = 8
	blackSpacesAmount = 32
)

// Colors
var black = color.RGBA{65, 30, 10, 255}
var white = color.RGBA{175, 140, 120, 255}

func main() {
	// create image filled with white color
	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))
	fillRect(img, img.Bounds(), white)

	// draw a chess board
	spaceSize := int(imageSize / boardSise)
	x, y := 0, 0
	for i := 0; i < blackSpacesAmount; i++ {
		fillRect(img, image.Rect(x, y, x+spaceSize, y+spaceSize), black)

		x += spaceSize * 2
		if x >= imageSize {
			x = 0
			y += spaceSize
			if (y/spaceSize)%2 != 0 {
				x = spaceSize
			}
		}
	}

	// save image to file
	if err := saveImageToFile(img, "rectangle.png"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Image has been created successfully")
	}
}
