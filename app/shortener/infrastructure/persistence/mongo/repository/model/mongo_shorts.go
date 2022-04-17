package model

type MongoShort struct {
	ShortUrl    string `json:"short_url" bson:"short_url,omitempty"`
	OriginalUrl string `json:"original_url" bson:"original_url,omitempty"`
	IsEnable    *bool  `json:"is_enable" bson:"is_enable,omitempty"`
	User        string `json:"user" bson:"user,omitempty"`
}
