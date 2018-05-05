package images

import (
	"fmt"
)

const UserRepo = "angelbarrera92/hasselhoffme"
const RawRepo = "https://raw.githubusercontent.com"
const FoldersRepo = "master/wallpapers"

func SearchGithubRawImages(w string) (result string) {
	url := fmt.Sprintf("%s/%s/%s/", RawRepo, UserRepo, FoldersRepo)

	return fmt.Sprintf("%s%d.jpg", url, RandomNumberInt(1, 9))
}
