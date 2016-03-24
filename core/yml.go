package core

import (
	"encoding/xml"
	"bytes"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
)

type YML struct {
	Date       string `xml:"date,attr"`
	Name       string `xml:"shop>name"`
	Company    string `xml:"shop>company"`
	Url        string `xml:"shop>url"`
	Categories []Categories `xml:"shop>categories>category"`
	Offers     []Offers `xml:"shop>offers>offer"`
}

type Categories struct {
	Name     string `xml:",innerxml"`
	Id       string `xml:"id,attr"`
	ParentId string `xml:"parentId,attr"`
}

type Offers struct {
	Id        int `xml:"id,attr"`
	Available string `xml:"available,attr"`
	Value     []NameField `xml:",any"`
}

type NameField struct {
	XMLName xml.Name `xml:""`
	Value   string   `xml:",chardata"`
}

func GetYML(data []byte) (YML, error) {
	yml := YML{}
	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReader
	err := decoder.Decode(&yml)
	if err != nil {
		return yml, err
	}
	return yml, nil
}