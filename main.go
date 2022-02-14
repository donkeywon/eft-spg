package main

import (
	"eft-spg/util"
	"fmt"
	"github.com/gobuffalo/packd"
)

func main() {
	util.ConfigFileBox.Walk(func(s string, file packd.File) error {
		fmt.Println(s)
		fmt.Println(file.Name())
		return nil
	})
}
