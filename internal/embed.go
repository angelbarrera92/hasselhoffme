package internal

import (
	"embed"
)

//go:embed wallpapers
var wallpapers embed.FS

func EmbededFiles() ([]string, error) {
	des, err := wallpapers.ReadDir("wallpapers")
	if err != nil {
		return nil, err
	}
	var files []string
	for _, de := range des {
		files = append(files, de.Name())
	}
	return files, nil
}

func EmbedFileContent(file string) ([]byte, error) {
	return wallpapers.ReadFile("wallpapers/" + file)
}
