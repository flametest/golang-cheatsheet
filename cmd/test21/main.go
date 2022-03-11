package main

import (
	"encoding/json"
	"errors"
	"fmt"
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

type Branch struct {
	branchCity      string
	branchSelection string
}

func NewBranch(branchCity, branchSelection string) *Branch {
	return &Branch{
		branchCity:      branchCity,
		branchSelection: branchSelection,
	}
}

func (b *Branch) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		BranchCity      string `json:"branch_city"`
		BranchSelection string `json:"branch_selection"`
	}{
		BranchCity:      b.branchCity,
		BranchSelection: b.branchSelection,
	})
}

func (b *Branch) UnmarshalJSON(bytes []byte) error {
	var tmp struct {
		BranchCity      string `json:"branch_city"`
		BranchSelection string `json:"branch_selection"`
	}
	err := json.Unmarshal(bytes, &tmp)
	if err != nil {
		return err
	}
	b.branchCity = tmp.BranchCity
	b.branchSelection = tmp.BranchSelection
	return nil
}


type ActivationForm struct {
	frontImages Photos
	backImages  Photos
	branch      *Branch
}

func NewActivationForm(frontImages Photos, backImages Photos, branch *Branch) (*ActivationForm, error) {
	if len(frontImages) == 0 {
		return nil, errors.New("front images must be provided")
	}
	return &ActivationForm{frontImages: frontImages, backImages: backImages, branch: branch}, nil
}

func (af *ActivationForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FrontImages Photos  `json:"front_images"`
		BackImages  Photos  `json:"back_images"`
		Branch      *Branch `json:"branch"`
	}{
		FrontImages: af.frontImages,
		BackImages:  af.backImages,
		Branch:      af.branch,
	})
}

func (af *ActivationForm) UnmarshalJSON(bytes []byte) error {
	var tmp struct {
		FrontImages Photos  `json:"front_images"`
		BackImages  Photos  `json:"back_images"`
		Branch      *Branch `json:"branch"`
	}
	err := json.Unmarshal(bytes, &tmp)
	if err != nil {
		return err
	}
	af.frontImages = tmp.FrontImages
	af.backImages = tmp.BackImages
	af.branch = tmp.Branch
	return nil
}

func (af *ActivationForm) FrontImages() Photos {
	return af.frontImages
}

func (af *ActivationForm) BackImages() Photos {
	return af.backImages
}

func main() {
	af, err := NewActivationForm([]*Photo{NewPhoto("fxx")} , NewPhotos([]string{"bxx"}), NewBranch("c", "s"))
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(af)
	marshal, err := af.MarshalJSON()
	if err != nil {
		fmt.Println(marshal)
	}
	fmt.Println(string(marshal))
	a := &ActivationForm{}
	//err = json.Unmarshal(marshal, a)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(a.frontImages[0].Url()	)
	err = a.UnmarshalJSON(marshal)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.frontImages[0].url)
}
