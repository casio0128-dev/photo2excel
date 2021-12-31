package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
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

/*
TODO: 処理の順序を明確に定義
TODO: 正式に実装
TODO: 署名データ作成用のコードを実装？
*/
func mask(src image.Image) {
	var img draw.Image

	//img = image.NewRGBA(image.Rect(0, 0, 500, 500))
	img = image.NewRGBA(src.Bounds())

	draw.Draw(img, src.Bounds(), src, image.Point{
		X: 0,
		Y: 0,
	}, draw.Over)

	m, _ := os.Open("image.png")
	mask, _ := png.Decode(m)

	for y := 0; y < 255; y++ {
		for x := 0; x < 255; x++ {
			r, g, b, a := mask.At(x, y).RGBA()

			img.Set(x, y, color.NRGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: uint8(a),
			})

			fmt.Println(img.At(x, y))
		}
	}

	f, _ := os.Create("sample2.png")

	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		log.Fatalln(err)
	}
}

func outputExample() {
	const width, height = 256, 256 // Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{R: uint8((x + y) & 255), G: uint8((x + y) << 1 & 255), B: uint8((x + y) << 2 & 255), A: 255})
		}
	}
	f, err := os.Create("image.png")
	defer func(fp *os.File) {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}(f)

	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
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
