package check

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
	"github.com/kirre02/monitor-backend/internal/site/service"
	"golang.org/x/sync/errgroup"
)

type Service struct {
	DB *sqlx.DB
}

// Function to initialize the Service with a valid DB connection
func NewCheckService(db *sqlx.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) check(ctx context.Context, site *service.Site) (*PingResponse, error) {
	// Perform a ping check on the site
	result, err := Ping(ctx, site.Url)
	if err != nil {
		return nil, err
	}

	// Insert the result into the database
	_, err = s.DB.ExecContext(ctx, `
        INSERT INTO checks (site_id, up, checked_at)
        VALUES ($1, $2, NOW())
    `, site.Id, result.Up)

	log.Infof("Checking site ID: %d", site.Id)

	return result, err
}

// Check checks a single site.
func (s *Service) Check(ctx context.Context, siteID int) (*PingResponse, error) {
	// Initialize Site Service
	siteSvc := service.NewSiteService(s.DB)

	// Retrieve the site that the user wants to check on
	site, err := siteSvc.Get(ctx, siteID)
	if err != nil {
		log.Error("Error checking site ID", siteID, ":", err)
		return nil, err
	}
	log.Infof("Checking site ID: %d", site.Id)

	return s.check(ctx, site)
}

func (s *Service) CheckAll(ctx context.Context) error {
	//initialize Site Service
	siteSvc := service.NewSiteService(s.DB)
	// Get all the tracked sites
	resp, err := siteSvc.List(ctx)
	if err != nil {
		log.Error("Error getting site list:", err)
		return err
	}

	// Check up to 8 sites concurrently.
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(8)
	for _, site := range resp.Sites {
		site := site
		g.Go(func() error {
			log.Infof("Checking URL: %s", site.Url)
			_, checkErr := s.check(ctx, site)
			return checkErr
		})
	}
	return g.Wait()
}

// TODO: add a cron job that will use the CheckAll function every hour or so
