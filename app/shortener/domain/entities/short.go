package entities

type Short struct {
	ShortUrl    string `json:"short_url"`
	OriginalUrl string `json:"original_url"`
	IsEnable    *bool  `json:"is_enable"`
	User        string `json:"user"`
}
