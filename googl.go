package googl

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/parnurzeal/gorequest"
)

type Googl struct {
	key string
}

type ShortMsg struct {
	Kind    string `json:"kind"`
	Id      string `json:"id"`
	LongUrl string `json:"longUrl"`
}

type LongMsg struct {
	Kind    string `json:"kind"`
	Id      string `json:"id"`
	LongUrl string `json:"longUrl"`
	Status  string `json:"status"`
}

func NewClient(key string) (*Googl, error) {
	if key == "" {
		return nil, fmt.Errorf("You need to set the Google Url Shortener API Key")
	}

	return &Googl{key: key}, nil
}

func (c *Googl) Shorten(url string) (link string, err error) {
	if url == "" {
		err = fmt.Errorf("You need to set the url to be shortened")
		return
	}

	request := gorequest.New()

	gUrl := "https://www.googleapis.com/urlshortener/v1/url?key=" + c.key

	resp, body, _ := request.Post(gUrl).
		Set("Accept", "application/json").
		Set("Content-Type", "application/json").
		Send(`{"longUrl":"` + url + `"}`).End()

	if resp.Status != "200 OK" {
		err = fmt.Errorf("Some error occurred, please try again later")
		return
	}

	decoder := json.NewDecoder(strings.NewReader(body))
	var b ShortMsg
	err = decoder.Decode(&b)
	if err != nil {
		err = fmt.Errorf("error decoding: %s", err)
		return
	}

	link = b.Id

	return link, nil
}
