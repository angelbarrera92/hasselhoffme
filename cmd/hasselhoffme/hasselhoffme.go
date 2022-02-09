package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/angelbarrera92/hasselhoffme/internal/images"
	"github.com/angelbarrera92/hasselhoffme/internal/motd"
	"github.com/reujab/wallpaper"
)

func main() {
	// Create CLI flags
	action := flag.String("action", "setwallpaper", "action to perform. Available actions: setwallpaper or setmotd")
	provider := flag.String("provider", "embeded", "image provider. Available providers: embeded, repository, bing")
	words := flag.String("words", "hasselhoff", "words to search for. Only working with bing provider")
	// Parse
	flag.Parse()

	// Default values
	fn := images.SearchEmbededImages
	remote := true

	// Validate
	switch *provider {
	case "embeded":
		fn = images.SearchEmbededImages
		remote = false
	case "repository":
		fn = images.SearchGithubRawImages
	case "bing":
		fn = images.SearchBingImage
		if *words == "" {
			fmt.Fprintf(os.Stderr, "provider bing requires words to search\n")
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown provider %s\n", *provider)
		os.Exit(1)
	}

	// Look for an image in the specified provider
	pathOrURL := searchRandomImage(fn, *words)
	if !remote {
		defer os.Remove(pathOrURL)
	}

	// Then, perform the specified action
	switch *action {
	case "setwallpaper":
		if remote {
			setWallpaperFromURL(pathOrURL)
		} else {
			setWallpaperFromPath(pathOrURL)
		}
	case "setmotd":
		if remote {
			setMotdFromURL(pathOrURL)
		} else {
			setMotdFromPath(pathOrURL)
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown action %s\n", *action)
		flag.PrintDefaults()
		os.Exit(1)
	}
}

// searchRandomImage searches for a random image using provided image search function (provider)
func searchRandomImage(sifn images.SearchImageFn, wordsToSearch string) string {
	return sifn(wordsToSearch)
}

func setWallpaperFromURL(url string) {
	err := wallpaper.SetFromURL(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while setting wallpaper from url %s: %v\n", url, err)
		os.Exit(1)
	}
}

func setWallpaperFromPath(path string) {
	err := wallpaper.SetFromFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while setting wallpaper from path %s: %v\n", path, err)
		os.Exit(1)
	}
}

func setMotdFromURL(url string) {
	data, err := download(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while downloading %s: %v\n", url, err)
		os.Exit(1)
	}
	motd.SetMotd(data)
}

func setMotdFromPath(path string) {
	data, err := readFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file %s: %v\n", path, err)
		os.Exit(1)
	}
	motd.SetMotd(data)
}

func download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while downloading %s: %v\n", url, err)
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

func readFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
