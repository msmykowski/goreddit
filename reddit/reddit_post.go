package reddit

import (
	"strings"
)

type RedditPost struct {
	Title string
	URL   string
	Score int
}

func (r RedditPost) GetImgurId() string {
	splitURL := strings.Split(r.URL, "/")
	imgurId := splitURL[3]

	if imgurId == "a" || imgurId == "g" || imgurId == "gallery" {
		imgurId = splitURL[4]
		return imgurId
	} else {
		imgurId = strings.Split(imgurId, ".")[0]
		return imgurId
	}
}
