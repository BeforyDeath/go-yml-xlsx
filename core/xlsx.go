package core

import (
	"github.com/tealeg/xlsx"
	"fmt"
)

func SetXlsx(yml YML) {
	xlsx_file, sheet := xlsx_init()

	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "id"
	cell = row.AddCell()
	cell.Value = "Available"

	for _, value := range yml.Offers[0].Value {
		cell := row.AddCell()
		cell.Value = value.XMLName.Local
	}

	for _, Offer := range yml.Offers {
		xlsx_add(sheet, &Offer)
	}

	xlsx_save(xlsx_file)
}

func xlsx_add(sheet *xlsx.Sheet, Offer *Offers) {
	row := sheet.AddRow()

	cell := row.AddCell()
	cell.SetInt(Offer.Id)
	cell = row.AddCell()
	cell.Value = Offer.Available

	for _, node := range Offer.Value {
		cell := row.AddCell()
		cell.Value = node.Value
	}
}

func xlsx_init() (file *xlsx.File, sheet *xlsx.Sheet) {
	file = xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf("xlsx error: create")
	}
	return
}

func xlsx_save(file *xlsx.File) {
	err := file.Save("tmp/xxxxxxxxxxx.xlsx")
	if err != nil {
		fmt.Printf("xlsx error: save")
	}
}
