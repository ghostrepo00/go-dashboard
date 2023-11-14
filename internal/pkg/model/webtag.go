package model

type Webtag struct {
	Id    string  `json:"id"`
	Url   string  `json:"url"`
	Title string  `json:"title"`
	Note  string  `json:"note"`
	Tags  []int64 `json:"tags"`
}
