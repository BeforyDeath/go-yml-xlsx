package core

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

func SetXlsx(yml YML, url string) (string, error) {
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

	return xlsx_save(xlsx_file, url)
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

func xlsx_save(file *xlsx.File, url string) (string, error) {

	hasher := md5.New()
	hasher.Write([]byte(url))
	name := hex.EncodeToString(hasher.Sum(nil))
	name = "tmp/" + name + ".xlsx"

	err := file.Save(name)
	if err != nil {
		return "", err
	}
	return name, nil
}
