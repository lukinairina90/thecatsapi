package catsapi

import (
	"fmt"
	"time"
)

type ListParams struct {
	Limit     int
	Page      int
	DescOrder bool
}

type ListResponse []CatImage

type CatImage struct {
	ID         string     `json:"id"`
	URL        string     `json:"url"`
	Categories []Category `json:"categories"`
	Breeds     []Breed    `json:"breeds"`
}

func (c CatImage) Info() string {
	return fmt.Sprintf("[ID] %s | [URL] %s | [CATEGORIES] %v | [BREEDS] %v\n", c.ID, c.URL, c.Categories, c.Breeds)
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Breed struct {
	AltNames       string `json:"alt_names"`
	Experimental   int    `json:"experimental"`
	Hairless       int    `json:"hairless"`
	Hypoallergenic int    `json:"hypoallergenic"`
	Id             string `json:"id"`
	LifeSpan       string `json:"life_span"`
	Name           string `json:"name"`
	Natural        int    `json:"natural"`
	Origin         string `json:"origin"`
	Rare           int    `json:"rare"`
	Rex            int    `json:"rex"`
	ShortLegs      int    `json:"short_legs"`
	SuppressedTail int    `json:"suppressed_tail"`
	Temperament    string `json:"temperament"`
	WeightImperial string `json:"weight_imperial"`
	WikipediaUrl   string `json:"wikipedia_url"`
}

type VoteBody struct {
	ImageId string `json:"image_id"`
	SubId   string `json:"sub_id"`
	Value   int    `json:"value"`
}

type ListRespVote []RespVote

type RespVote struct {
	Id          int       `json:"id"`
	ImageId     string    `json:"image_id"`
	SubId       string    `json:"sub_id"`
	CreatedAt   time.Time `json:"created_at"`
	Value       int       `json:"value"`
	CountryCode string    `json:"country_code"`
	Image       struct {
		Id  string `json:"id"`
		Url string `json:"url"`
	} `json:"image"`
}

func (v RespVote) Info() string {
	return fmt.Sprintf("[ID] %d | [IMAGEID] %s | [SUBID] %s | [CREATED_AT] %s | [VALUE] %d | [COUNTRY_CODE] %s | [IMAGE] %v",
		v.Id, v.ImageId, v.SubId, v.CreatedAt, v.Value, v.CountryCode, v.Image)
}

type Message struct {
	Message string `json:"message"`
}
