package model

import (
	"errors"
	"fmt"
	"regexp"
)

type CreateInput struct {
	Url  string `json:"url"`
	User string `json:"user"`
}

type CreateOutput struct {
	ShortUrl string `json:"short_url"`
}

type GetOutput struct {
	Url string `json:"url"`
}

type ModifyInput struct {
	OriginalUrl string `json:"original_url,omitempty"`
	IsEnable    *bool  `json:"is_enable,omitempty"`
	User        string `json:"user,omitempty"`
}

func (ci *CreateInput) ValidateUrl() error {
	regexpMatch := "(www|http://|https://)[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)"
	matched, _ := regexp.MatchString(regexpMatch, fmt.Sprintf("%s", ci.Url))
	if !matched {
		return errors.New("input is not a valid url")
	}
	return nil
}
