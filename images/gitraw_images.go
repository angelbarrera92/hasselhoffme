package images

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const UserRepo = "angelbarrera92/hasselhoffme"

type Links struct {
	Self string `json:"message"`
	Git  string `json:"git"`
	HTML string `json:"html"` // nolint
}

type Content struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HtmlURL     string `json:"html_url"`     // nolint
	GitURL      string `json:"git_url"`      // nolint
	DownloadURL string `json:"download_url"` // nolint
	TypeObject  string `json:"type"`
	Links       Links  `json:"links"`
}

func SearchGithubRawImages(w string) (result string) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/contents/wallpapers", UserRepo)

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
