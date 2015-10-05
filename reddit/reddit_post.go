package reddit

import ()

type RedditPost struct {
	Title string
	URL   string
	Score int
}

func (r RedditPost) GetImgurId() string {
	return r.URL
}
