package images

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

const UserRepo = "ervitis/hasselhoffme"
const RawRepo = "https://raw.githubusercontent.com"
const FoldersRepo = "master/wallpapers"

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
	url := fmt.Sprintf("%s/%s/%s/", RawRepo, UserRepo, FoldersRepo)

	v := getLastIndex(url)
	if v == 0 {
		panic("Could not load any image")
	}

	return fmt.Sprintf("%s%d.jpg", url, RandomNumberInt(1, v))
}

func getLastIndex(baseUrl string) int {
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/contents/wallpapers", UserRepo)

	var content []Content

	resp, _ := http.Get(apiUrl)
	if resp.StatusCode >= 400 {
		return 0
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0
	}

	err = json.Unmarshal(body, &content)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return len(content)
}
