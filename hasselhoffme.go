package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/angelbarrera92/hasselhoffme/images"
	"github.com/reujab/wallpaper"
	"github.com/zyxar/image2ascii/ascii"
)

const MOTD_FILE = "/etc/update-motd.d/99-hasselhoffme"

func usage() {
	fmt.Fprintf(os.Stderr, `%s [-h] [<action>]
  available actions:
	setwallpaper: sets a random wallpaper
			(default if no action specified)
	setmotd: sets a random ascii art motd
`, os.Args[0])
	os.Exit(1)
}

func main() {
	action := "setwallpaper"
	url := SearchRandomImage(images.SearchGithubRawImages, "")

	if len(os.Args) >= 2 {
		action = os.Args[1]
	}

	switch action {
	case "setwallpaper":
		setWallpaperFromURL(url)
	case "setmotd":
		setMotdFromURL(url)
	default:
		usage()
	}
}

func SearchRandomImage(sifn images.SearchImageFn, wordsToSearch string) string {
	return sifn(wordsToSearch)
}

func setWallpaperFromURL(url string) {
	wallpaper.SetFromURL(url)
}

func setMotdFromURL(url string) {
	if os.Getuid() != 0 {
		fmt.Fprintf(os.Stderr, "motd can only be changed as root\n")
		os.Exit(1)
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while downloading %s: %v\n", url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	opt := ascii.Options{
		Width:  0,
		Height: 0,
		Color:  false,
		Invert: false,
		Flipx:  false,
		Flipy:  false}
	motd, err := ascii.Decode(resp.Body, opt)

	f, err := os.OpenFile(MOTD_FILE, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while opening %s: %v\n", MOTD_FILE, err)
		os.Exit(1)
	}
	defer f.Close()

	fmt.Fprintf(f, `#!/bin/sh
cat <<EOF
`)

	motd.WriteTo(f)
	fmt.Fprintf(f, "EOF\n")

}
