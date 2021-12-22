package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func getImageSize(img image.Image) (height, width int) {
	bound := img.Bounds()

	height = bound.Max.Y
	width = bound.Max.X
	return
}

func openPicture(path string) (image.Image, error) {
	if strings.EqualFold(path, BLANK) {
		return nil, fmt.Errorf("openPicture 1 Error")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func createPhotoRGBMap(path string) ([][]string, error) {
	img, err := openPicture(path)

	if err != nil {
		return nil, err
	}

	height, width := getImageSize(img)

	colorMap := [][]string{}

	for x := 0; x < width; x++ {
		colorRow := []string{}
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			rgbString := getRGBString(r, g, b)
			colorRow = append(colorRow, rgbString)
		}
		colorMap = append(colorMap, colorRow)
	}

	return colorMap, nil
}
