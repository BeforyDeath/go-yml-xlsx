package core

import (
	"os"
	"encoding/json"
)

type Configuration struct {
	BasePath string
}

var Config Configuration

func (c *Configuration) Get()  {
	LogInf.Println("Read config file")
	file, err := os.Open("config.json")
	if err != nil {
		LogErr.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		LogErr.Fatal(err)
	}
}
