package main

import (
	"path/filepath"

	"github.com/NickolayGoloburdin/tlg_bot_golang/internal/app"
)

var filename string = "config.json"

func main() {
	a := app.NewApp(filepath.Join("../", "config.json"))
	a.Start()
}
