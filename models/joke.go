package models

type Joke struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

type JokeRequest struct {
	Num int `jsoon:"num"`
}
