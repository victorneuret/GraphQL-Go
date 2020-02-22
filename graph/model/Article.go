package model

type Article struct {
	ID   string `json:"id"`
	Title string `json:"text"`
	Text string `json:"text"`
	UserID string  `json:"user"`
}

type NewArticle struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID string `json:"userId"`
}