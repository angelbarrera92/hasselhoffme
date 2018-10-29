package images

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// UserRepo is the name of GitHub repository with wallpaper images
const UserRepo = "angelbarrera92/hasselhoffme"

// Links represents the `links` section in GitHub JSON API response
type Links struct {
	Self string `json:"message"`
	Git  string `json:"git"`
	HTML string `json:"html"`
}

// Content represents the GitHub JSON API repositories/contents response body
type Content struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	GitURL      string `json:"git_url"`
	DownloadURL string `json:"download_url"`
	TypeObject  string `jbson:"type"`
	Links       Links  `json:"links"`
}

// SearchGithubRawImages returns a random image result
func SearchGithubRawImages(w string) (result string) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/contents/internal/wallpapers", UserRepo)

	content, err := getContent(url)
	if err != nil {
		panic("Could not load any image")
	}

	var images []ImageResult

	for k, v := range content {
		images = append(images, ImageResult{
			Source: v.DownloadURL,
			Index:  k,
		})
	}

	rn := RandomNumber(images)

	return images[rn].Source
}

func getContent(baseURL string) ([]Content, error) {

	var content []Content

	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error response: %s", resp.Body)
	}

	defer resp.Body.Close()

	body, ioerr := ioutil.ReadAll(resp.Body)
	if ioerr != nil {
		return nil, ioerr
	}

	err = json.Unmarshal(body, &content)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return content, nil
}
