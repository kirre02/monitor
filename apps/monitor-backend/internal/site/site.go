package site

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
)

type Site struct {
	// the unqie id for the site
	Id int `json:"id"`
	// the Url of the site
	Url string `json:"Url"`
}

// The params for adding a site to be monitored
type AddParams struct {
	// Url is the Url of the site that we want to be monitored
	Url string `json:"Url"`
}

type Service struct {
	DB     *sqlx.DB
	Logger log.Logger
}

func NewSiteService(db *sqlx.DB, logger log.Logger) *Service {
	return &Service{
		DB:     db,
		Logger: logger,
	}
}

func (s *Service) Add(ctx context.Context, p *AddParams) (*Site, error) {
	s.Logger.Info("Adding website...")

	site := &Site{Url: p.Url}
	_, err := s.DB.NamedExecContext(ctx, "INSERT INTO sites (Url) VALUES (:Url)", site)
	if err != nil {
		s.Logger.Error("Failed to add website", "error", err)
		return nil, err
	}

	s.Logger.Info("Website was successfully added")
	return site, nil
}

func (s *Service) Get(ctx context.Context, siteID int) (*Site, error) {
	var site Site

	s.Logger.Info("Fetching website...")
	err := s.DB.GetContext(ctx, &site, "SELECT id, Url FROM sites WHERE id = $1", siteID)
	if err != nil {
		s.Logger.Error("Failed fetching website", err)
		return nil, err
	}

	s.Logger.Info("Website successfully retrieved")
	return &site, nil
}

func (s *Service) Delete(ctx context.Context, SiteID int) error {
	_, err := s.DB.ExecContext(ctx, "DELETE FROM sites WHERE id = $1", SiteID)

	s.Logger.Info("Website successfully deleted")
	return err
}

type ListResponse struct {
	//sites is the list of monitored sites.
	Sites []*Site `json:sites`
}

func (s *Service) List(ctx context.Context) (*ListResponse, error) {
	var sites []*Site

	s.Logger.Info("Fetching website list...")
	err := s.DB.SelectContext(ctx, &sites, "SELECT id, Url FROM sites")
	if err != nil {
		s.Logger.Error("Failed to fetch site list", "error", err)
		return nil, err
	}

	s.Logger.Info("Site list retrieved successfully")
	return &ListResponse{Sites: sites}, nil
}
