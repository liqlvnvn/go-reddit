package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Item struct {
	Title    string
	URL      string
	Comments int `json:"num_comments"`
}

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

func (i Item) String() string {
	com := ""
	switch i.Comments {
	case 0:
		// nothing
	case 1:
		com = " (1 comment"
	default:
		com = fmt.Sprintf(" (%d comments)", i.Comments)
	}
	return fmt.Sprintf("%s%s\n%s", i.Title, com, i.URL)
}

const UserAgent = "Golang Reddit Reader"

func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)

	// Create a request and add the proper headers.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)

	// Handle the request
	// resp, err := http.Get(url)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Data
	}
	return items, nil
}

func main() {
	items, err := Get("goland")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item.Title)
	}
}
