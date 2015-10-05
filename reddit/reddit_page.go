package reddit

import (
	"encoding/json"
	"net/http"
)

type RedditPage struct {
	Data struct {
		Children []struct {
			Data RedditPost
		}
	}
}

const redditUrl = "http://reddit.com/"

func GetRedditPage() *RedditPage {
	resp, _ := http.Get(redditUrl + ".json")
	defer resp.Body.Close()

	r := new(RedditPage)

	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(r)

	return r
}

func GetSubredditPage(subreddit string) *RedditPage {
	resp, _ := http.Get(redditUrl + "/r/" + subreddit + ".json")
	defer resp.Body.Close()

	r := new(RedditPage)

	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(r)

	return r
}
