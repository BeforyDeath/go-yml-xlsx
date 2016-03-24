package main

import (
	"github.com/beforydeath/go-yml-xlsx/core"
)

func main() {
	defer core.LogClose()
	core.Config.Get()
}