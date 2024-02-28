package check

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
	"github.com/kirre02/monitor-backend/internal/site/service"
	"github.com/robfig/cron"
	"golang.org/x/sync/errgroup"
)

type Service struct {
	DB *sqlx.DB

	Cron *cron.Cron
}

// Function to initialize the Service with a valid DB connection
func NewCheckService(db *sqlx.DB) (*Service, error) {
	checkService := &Service{DB: db, Cron: cron.New()}

	err := checkService.Cron.AddFunc("0 */30 * * *", func() {
		err := checkService.CheckAll(context.Background())
		if err != nil {
			log.Errorf("difficulty running CheckAll, %s", err)
		}
	})
	if err != nil {
		return nil, err
	}

	checkService.Cron.Start()

	return checkService, nil
}

func (s *Service) StopCron() {
	if s.Cron != nil {
		s.Cron.Stop()
	}
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
	log.Infof("Running CheckAll at: %s", time.Now().Format(time.RFC3339))
	// initialize Site Service
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
