package main

import (
	"github.com/reujab/wallpaper"
)

func main() {
	url := SearchHasselhoffRandom("david hasselhoff")
	setWallpaperFromURL(url)
}

func setWallpaperFromURL(url string) {
	wallpaper.SetFromURL(url)
}

