package main

import (
	"github.com/alecthomas/kingpin/v2"
	_ "image/jpeg"
	"log"
	"path/filepath"
	"strings"
	"sync"
)

/*
	TODO: 並列化ができるか試す、複数ブックで処理を始めて、最後にひとつにマージ？　別シートなら並列いける？
	TODO: 全体的二里ファクタ
*/

var (
	input        string
	output       string
	goRoutineMax int
)

func init() {
	kingpin.Flag("image", "Please specify the image you want to convert to Excel.").Short('i').Required().StringVar(&input)
	kingpin.Flag("output", "Please specify the output destination.").Short('o').Default(getFileName(input)).StringVar(&output)
	kingpin.Flag("activeRoutine", "Specify the number of coroutines to run. In addition, if it is too large, the load will be applied to the PC. Default 10 coroutines").Short('r').Default("10").IntVar(&goRoutineMax)
	if !strings.HasSuffix(filepath.Ext(output), ".xlsx") {
		output += ".xlsx"
	}
	kingpin.Parse()
}

func main() {
	colors, err := createPhotoRGBMap(input)
	if err != nil {
		log.Fatalln(err)
	}

	book, err := openExcelBook(output)
	if err != nil {
		log.Fatalln(err)
	}

	routineManager := make(chan struct{}, goRoutineMax)
	wg := &sync.WaitGroup{}
	for x, colorRow := range colors {
		routineManager <- struct{}{}
		x := x
		colorRow := colorRow
		wg.Add(1)
		go func(cr []string) {
			defer wg.Done()
			for y, color := range cr {
				y := y
				color := color
				wg.Add(1)
				go func() {
					defer wg.Done()
					fillCellColor(book, color, x+1, y+1)
				}()
			}
		}(colorRow)
	}
	wg.Wait()
	if err := book.SaveAs(output); err != nil {
		panic(err)
	}
}

func getFileName(target string) string {
	fileName := filepath.Base(target)
	fileExt := filepath.Ext(fileName)
	return strings.Split(fileName, fileExt)[0]
}
