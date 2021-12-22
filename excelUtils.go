package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"path/filepath"
)

func openExcelBook(path string) (*excelize.File, error) {
	var book *excelize.File
	var err error

	if !isFileExists(path) {
		book = excelize.NewFile()
		err = book.SaveAs(filepath.Base(path))
	} else {
		book, err = excelize.OpenFile(path)
	}

	if err != nil {
		return nil, err
	}

	return book, nil
}

func fillCellColor(book *excelize.File, rgb string, x, y int) {
	style, err := book.NewStyle(fmt.Sprintf(`{"fill":{"type":"pattern","color":["%s"],"pattern":1}}`, rgb))

	if err != nil {
		log.Fatalln(err)
	}

	columnName, err := excelize.ColumnNumberToName(x)
	if err != nil {
		log.Fatalln(err)
	}

	book.SetCellStyle("Sheet1", fmt.Sprintf("%s%d", columnName, y+1), fmt.Sprintf("%s%d", columnName, y+1), style)
}
