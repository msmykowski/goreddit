package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

type Item struct {
	Title string
	URL string
	Comments int `json:"num_comments"`
}


func main() {
	items := getReddit(redditUrl)
	for _, item := range items {
		fmt.Println(item)
	}
}

const redditUrl = "http://reddit.com/.json"

func getReddit(url string) []Item {
	resp, _ := http.Get(redditUrl)
	defer resp.Body.Close()
	
	r := new(Response)

	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(r)

	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Data
	}

	return items
}
