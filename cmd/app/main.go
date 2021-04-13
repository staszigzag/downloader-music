package main

import (
	"flag"

	"github.com/staszigzag/downloader-music/internal/app"
)

func main() {
	configPath := flag.String("configPath", "./config", "App config path")
	flag.Parse()

	app.Run(*configPath)
}
