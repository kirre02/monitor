package service

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Site struct {
	// the unqie id for the site
	Id int `json:"id"`
	// the Url of the site
	Url string `json:"Url"`
	// name of the site
	Name string `json:"Name"`
}

// The params for adding a site to be monitored
type AddParams struct {
	// Url is the Url of the site that we want to be monitored
	Url string `json:"Url"`
	// Name of the site that user wants to call it
	Name string `json:"Name"`
}

type SiteServiceInterface interface {
	Add(ctx context.Context, p *AddParams) (*Site, error)
	Get(ctx context.Context, siteID int) (*Site, error)
	List(ctx context.Context) (*ListResponse, error)
	Delete(ctx context.Context, SiteID int) error
}

type Service struct {
	DB *sqlx.DB
}

// Function to initialize the Service with a valid DB connection
func NewSiteService(db *sqlx.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) Add(ctx context.Context, p *AddParams) (*Site, error) {
	// Check if AddParams pointer is not nil
	if p == nil {
		return nil, fmt.Errorf("AddParams cannot be nil")
	}

	// Check if URL or Name is empty in AddParams
	if p.Url == "" || p.Name == "" {
		return nil, fmt.Errorf("URL and Name cannot be empty")
	}

	site := &Site{Url: p.Url, Name: p.Name}

	if s.DB == nil {
		return nil, fmt.Errorf("DB connection is nil")
	}

	_, err := s.DB.NamedExecContext(ctx, "INSERT INTO sites (url, name) VALUES (:url, :name)", site)
	if err != nil {
		return nil, err
	}

	return site, nil
}

func (s *Service) Get(ctx context.Context, siteID int) (*Site, error) {
	var site Site

	err := s.DB.GetContext(ctx, &site, "SELECT id, url, name FROM sites WHERE id = $1", siteID)
	if err != nil {
		return nil, err
	}

	return &site, nil
}

func (s *Service) Delete(ctx context.Context, SiteID int) error {
	_, err := s.DB.ExecContext(ctx, "DELETE FROM sites WHERE id = $1", SiteID)

	return err
}

type ListResponse struct {
	// sites is the list of monitored sites.
	Sites []*Site `json:"sites"`
}

func (s *Service) List(ctx context.Context) (*ListResponse, error) {
	var sites []*Site

	err := s.DB.SelectContext(ctx, &sites, "SELECT id, Url, Name FROM sites")
	if err != nil {
		return nil, err
	}

	return &ListResponse{Sites: sites}, nil
}
