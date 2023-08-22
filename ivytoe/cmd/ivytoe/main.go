package main

import (
	"example/taylormonacelli/ivytoe"
	"os"
)

func main() {
	status := ivytoe.Main()
	os.Exit(status)
}
