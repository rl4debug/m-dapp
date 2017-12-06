package dapp

type ProjectItem struct {
	Id          string `json:"id"`
	Badges      []string
	Author      string  `json:"author"`
	SubAuthors  string  `json:"additionalAuthors,omitempty"`
	IsNew       bool    `json:"isNew"`
	IsNsfw      bool    `json:"isNsfw"`
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Status      string  `json:"status"`
	Teaser      string  `json:"teaser"`
	Impressions int     `json:"impressions"`
	Clicks      int     `json:"clicks"`
	Score       float64 `json:"score"`
}

type ProjectItems []ProjectItem

type ProjectDetailItem struct {
	Des  string   `json:"description"`
	Tags []string `json:tags`
	Site string   `json:"website"`
}
