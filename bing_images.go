package main

import (
	url2 "net/url"
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/moovweb/gokogiri"
	"math/rand"
	"time"
)

type ImageResult struct {
	Source string
	Index int
}

const BaseUrl = "http://www.bing.com/images/search?q="

func SearchHasselhoffRandom(searchWord string) (result string) {
	page, err := fetchPage(searchWord); if err != nil {
		return err.Error()
	}

	images, err := parseResult(page); if err != nil {
		return err.Error()
	}

	rn := randomNumber(images)

	return images[rn].Source
}

func fetchPage(searchWord string) (results []byte, err error) {
	url := BaseUrl + url2.QueryEscape(searchWord)

	resp, err := http.Get(url); if err != nil {
		return nil, err
	} else if resp.StatusCode > 203 {
		return nil, fmt.Errorf("response code %d", resp.StatusCode)
	}

	page, err := ioutil.ReadAll(resp.Body); if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return page, nil
}

func parseResult(html []byte) (results []ImageResult, err error) {
	body, err := gokogiri.ParseHtml([]byte(html)); if err != nil {
		return nil, err
	}

	root := body.Root()
	previews, err := root.Search("//a/div/img"); if err != nil {
		return nil, err
	}

	var images []ImageResult
	c := 0
	for _, v := range previews {
		src := v.Attr("src")
		images = append(images, ImageResult{
			Source: src,
			Index: c,
		})
		c ++
	}

	return images, nil
}

func randomNumber(images []ImageResult) int {
	rand.Seed(time.Now().Unix())

	return rand.Intn(len(images))
}