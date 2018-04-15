package main

import (
	"github.com/reujab/wallpaper"
)

func main() {
	url := "https://raw.githubusercontent.com/angelbarrera92/hasselhoffme/master/wallpaper.jpg"
	setWallpaperFromURL(url)
}

func setWallpaperFromURL(url string) {
	wallpaper.SetFromURL(url)
}

