package images

// ImageResult represents a single search result along with its number in results list
type ImageResult struct {
	Source string
	Index  int
}

// SearchImageFn that returns a random image search result for given query
type SearchImageFn func(w string) string
