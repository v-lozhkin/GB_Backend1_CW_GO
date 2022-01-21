package models

import (
	"errors"
	"net/url"
)

type Link struct {
	ID   int    `json:"-"`
	Link string `json:"link"`
}

func (l Link) Validate() error {
	if l.Link == "" {
		return errors.New("link can't be empty")
	}

	if _, err := url.ParseRequestURI(l.Link); err != nil {
		return errors.New("link is invalid")
	}

	return nil
}
