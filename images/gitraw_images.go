package images

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

const UserRepo = "ervitis/hasselhoffme"

type Links struct {
	Self string `json:"message"`
	Git string `json:"git"`
	Html string `json:"html"`
}

type Content struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Sha string `json:"sha"`
	Size int `json:"size"`
	Url string `json:"url"`
	HtmlUrl string `json:"html_url"`
	GitUrl string `json:"git_url"`
	DownloadUrl string `json:"download_url"`
	TypeObject string `json:"type"`
	Links Links `json:"links"`
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
			Source: v.DownloadUrl,
			Index: k,
		})
	}

	rn := RandomNumber(images)

	return images[rn].Source
}

func getContent(baseUrl string) ([]Content, error) {

	var content []Content

	resp, _ := http.Get(baseUrl)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error response: %s", resp.Body)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &content)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return content, nil
}
