package reddit

import (
	"strings"
)

type RedditPost struct {
	Title string
	URL   string
	Score int
}

func (r RedditPost) GetImgurId() (i string) {
	splitURL := strings.Split(r.URL, "/")
	indicator := splitURL[3]

	if indicator == "a" || indicator == "g" || indicator == "gallery" {
		i = splitURL[4]
	} else if strings.Contains(indicator, ".") {
		i = strings.Split(indicator, ".")[0]
	} else {
		i = indicator
	}

	return i
}
