package site

import (
	"context"

	"github.com/jmoiron/sqlx"

)

type Site struct {
	// the unqie id for the site
	id int `json:"id"`
	// the url of the site
	url string `json:"url"`
}

// The params for adding a site to be monitored
type AddParams struct {
	// url is the URL of the site that we want to be monitored
	url string `json:"url"`
}

type Service struct {
	db *sqlx.DB
}

func (s *Service) Add(ctx context.Context, p *AddParams) (*Site, error) {
	site := &Site{url: p.url}
	return site, nil
}



