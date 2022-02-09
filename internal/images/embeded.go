package images

import (
	"github.com/angelbarrera92/hasselhoffme/internal"
)

// SearchEmbededImages returns a random image result
func SearchEmbededImages(w string) (result string) {
	files, err := internal.EmbededFiles()
	if err != nil {
		return ""
	}

	var images []ImageResult
	for k, v := range files {
		images = append(images, ImageResult{
			Source: v,
			Index:  k,
		})
	}
	rn := RandomNumber(images)

	wallPaperContent, err := internal.EmbedFileContent(images[rn].Source)
	if err != nil {
		return ""
	}

	f, err := createTempFile(wallPaperContent)
	if err != nil {
		return ""
	}

	return f.Name()
}
