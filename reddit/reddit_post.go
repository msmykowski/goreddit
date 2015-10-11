package reddit

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strings"
)

type RedditPost struct {
	Title string
	URL   string
	Score int
}

type ImgurPost struct {
	Data ImgurData
}

type ImgurData struct {
	Id string
	Link string
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

func (r RedditPost) GetImgurUrl(id string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.imgur.com/3/image/" + id , nil)
	req.Header.Add("Authorization", "Client-ID 95abf06f166b929")
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	
	i := new(ImgurPost)
	
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(i)
	
	fmt.Println(i)
}
