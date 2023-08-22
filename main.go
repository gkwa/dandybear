package main

import (
	"example/taylormonacelli/dandybear/ivytoe"
)

var logger *ivytoe.Logger

func init() {
	logger = ivytoe.NewLogger("dandybear.log")
}

func main() {
	logger.Fatal("test")
}
