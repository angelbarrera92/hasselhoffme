package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/angelbarrera92/hasselhoffme/internal/images"
	"github.com/reujab/wallpaper"
	"github.com/zyxar/image2ascii/ascii"
)

const (
	// MOTDFile is a path to /etc/motd file
	MOTDFile = "/etc/motd"
	// UpdateMOTDPath is a path to /etc/update-motd.d
	UpdateMOTDPath = "/etc/update-motd.d"
	// UpdateMOTDFile is a path to /etc/update-motd.d/99-hasselhofme
	UpdateMOTDFile = UpdateMOTDPath + "/99-hasselhoffme"
)

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

// SearchRandomImage searches for a random image using provided image search function
func SearchRandomImage(sifn images.SearchImageFn, wordsToSearch string) string {
	return sifn(wordsToSearch)
}

func setWallpaperFromURL(url string) {
	err := wallpaper.SetFromURL(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while setting wallpaper from url %s: %v\n", url, err)
		os.Exit(1)
	}
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
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while decoding url response body %s: %v\n", url, err)
		os.Exit(1)
	}

	if _, err := os.Stat(UpdateMOTDPath); os.IsNotExist(err) {
		writeMotd(motd)
	} else {
		writeUpdateMotdScript(motd)
	}
}

func writeMotd(motd *ascii.Ascii) {
	content := ""
	if _, err := os.Stat(MOTDFile); !os.IsNotExist(err) {
		contentBytes, err := ioutil.ReadFile(MOTDFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while opening %s: %v\n", MOTDFile, err)
			os.Exit(1)
		}
		content = string(contentBytes)
	}

	re := regexp.MustCompile(`(?s)### hasselhon ###.*### hasselhoff ###\n`)
	content = re.ReplaceAllString(content, "")

	f, err := os.OpenFile(MOTDFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755) // nolint
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while opening %s: %v\n", MOTDFile, err)
		os.Exit(1)
	}
	defer f.Close()

	fmt.Fprintf(f, "%s### hasselhon ###\n", content)
	_, err = motd.WriteTo(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while writing to file %s: %v\n", f.Name(), err)
		os.Exit(1)
	}
	fmt.Fprintf(f, "### hasselhoff ###\n")
}

func writeUpdateMotdScript(motd *ascii.Ascii) {
	f, err := os.OpenFile(UpdateMOTDFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755) // nolint
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while opening %s: %v\n", UpdateMOTDFile, err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, `#!/bin/sh
cat <<EOF
`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while outputting to file %s: %v\n", f.Name(), err)
		os.Exit(1)
	}

	_, err = motd.WriteTo(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while writing to file %s: %v\n", f.Name(), err)
		os.Exit(1)
	}

	_, err = fmt.Fprintf(f, "EOF\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while outputting to file %s: %v\n", f.Name(), err)
		os.Exit(1)
	}
}
