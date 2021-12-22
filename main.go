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

	fmt.Println(colors[0][0])
}
