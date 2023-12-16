package models

type URL struct{
	ShortURL string `json:"short_url" bson:"short_url"`
    LongURL  string `json:"long_url" bson:"long_url"`
}