package images

type ImageResult struct {
	Source string
	Index  int
}

type SearchImageFn func(w string) string
