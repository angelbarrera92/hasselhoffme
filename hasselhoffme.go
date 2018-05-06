package main

import (
	"github.com/reujab/wallpaper"
	"github.com/angelbarrera92/hasselhoffme/images"
)

func main() {
	url := SearchRandomImage(images.SearchGithubRawImages, "")

	setWallpaperFromURL(url)
}

func SearchRandomImage(sifn images.SearchImageFn, wordsToSearch string) string {
	return sifn(wordsToSearch)
}

func setWallpaperFromURL(url string) {
	wallpaper.SetFromURL(url)
}

