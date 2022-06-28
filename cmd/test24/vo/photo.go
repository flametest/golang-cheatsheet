package vo

import (
	"encoding/json"
)

type Photo struct {
	url string
}

func (p *Photo) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Url string `json:"url"`
	}{
		Url: p.url,
	})
}

func (p *Photo) UnmarshalJSON(bytes []byte) error {
	var tmp struct {
		Url string `json:"url"`
	}
	err := json.Unmarshal(bytes, &tmp)
	if err != nil {
		return err
	}
	p.url = tmp.Url
	return nil
}

func (p *Photo) Url() string {
	return p.url
}

func NewPhoto(url string) *Photo {
	return &Photo{url: url}
}

func NewPhotos(urls []string) Photos {
	photos := make(Photos, 0)
	for _, url := range urls {
		photos = append(photos, NewPhoto(url))
	}
	return photos
}

type Photos []*Photo

func (p Photos) GetUrls() []string {
	urls := make([]string, 0)
	for _, photo := range p {
		urls = append(urls, photo.Url())
	}
	return urls
}

func (p Photos) GetOne() *Photo {
	if len(p) != 0 {
		return p[0]
	}
	return nil
}

func EmptyPhotos() Photos {
	return make(Photos, 0)
}
