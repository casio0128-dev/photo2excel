package main

import (
	"fmt"
	_ "image/jpeg"
	"log"
)

/*
	TODO: 並列化ができるか試す、複数ブックで処理を始めて、最後にひとつにマージ？　別シートなら並列いける？
	TODO: 全体的二里ファクタ
*/

func main() {
	colors, err := createPhotoRGBMap("./sample.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	book, err := openExcelBook("sample.xlsx")
	if err != nil {
		log.Fatalln(err)
	}

	for x, colorRow := range colors {
		for y, color := range colorRow {
			fillCellColor(book, color, x+1, y+1)
		}
	}

	book.SaveAs("sample.xlsx")

	fmt.Println(colors[0][0])
}
