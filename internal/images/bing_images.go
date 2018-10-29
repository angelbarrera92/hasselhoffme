package images

import (
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strings"

	"golang.org/x/net/html"
)

// BaseURL is the Bing Image Search endpoint
const BaseURL = "http://www.bing.com/images/search?q="

// SearchBingImage returns a random Bing Image Search result
func SearchBingImage(searchWord string) (result string) {
	images, err := parseResult(searchWord)
	if err != nil {
		return err.Error()
	}

	rn := RandomNumber(images)

	return images[rn].Source
}

func parseResult(searchWord string) (results []ImageResult, err error) {
	url := BaseURL + url2.QueryEscape(searchWord)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	} else if resp.StatusCode > 203 {
		return nil, fmt.Errorf("response code %d", resp.StatusCode)
	}

	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyHTML := strings.Replace(string(page), "<noscript>", "", -1)
	bodyHTML = strings.Replace(bodyHTML, "</noscript>", "", -1)

	var images []ImageResult

	if document, err := html.Parse(strings.NewReader(bodyHTML)); err == nil {

		var parser func(node *html.Node)

		parser = func(node *html.Node) {
			if node.Type == html.ElementNode && node.Data == "img" {
				c := 0

				for _, e := range node.Attr {
					if e.Key == "src" && strings.Contains(e.Val, "http") {
						images = append(images, ImageResult{
							Source: e.Val,
							Index:  c,
						})
						c++
					}
				}
			}

			for el := node.FirstChild; el != nil; el = el.NextSibling {
				parser(el)
			}

		}

		parser(document)
	}

	return images, nil
}
