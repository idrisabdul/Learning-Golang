package utils

import (
	"bytes"

	"github.com/tealeg/xlsx"
)

func ExportExcell(data [][]string) []byte {
	// Membuat file Excel baru
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		panic(err)
	}

	// Menulis data ke sheet
	for _, row := range data {
		rowCell := sheet.AddRow()
		for _, cell := range row {
			rowCell.AddCell().SetValue(cell)
		}
	}

	buffer := new(bytes.Buffer)
	file.Write(buffer)
	return buffer.Bytes()
}

func ExportExcellCustom(data [][]string, header string) []byte {
	// Membuat file Excel baru
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		panic(err)
	}

	// membuat style bold
	style := xlsx.NewStyle()
	style.Font.Bold = true

	// Menambahkan baris untuk header dan mengatur style bold
	headerRow := sheet.AddRow()
	headerCell := headerRow.AddCell()
	headerCell.SetValue(header)
	headerCell.SetStyle(style)

	// Menulis data ke sheet
	for _, row := range data {
		rowCell := sheet.AddRow()
		for _, cell := range row {
			rowCell.AddCell().SetValue(cell)
		}
	}

	buffer := new(bytes.Buffer)
	file.Write(buffer)
	return buffer.Bytes()
}
