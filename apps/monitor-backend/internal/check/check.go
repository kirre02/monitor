package check

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
	"github.com/kirre02/monitor-backend/internal/site/service"
	"golang.org/x/sync/errgroup"
)

type SiteService struct {
	Site service.SiteServiceInterface
}

func check(ctx context.Context, site *service.Site, db *sqlx.DB) error {
	// Perform a ping check on the site
	result, err := Ping(ctx, site.Url)
	if err != nil {
		return err
	}

	// Insert the result into the database
	_, err = db.ExecContext(ctx, `
		INSERT INTO checks (site_id, up, checked_at)
		VALUES ($1, $2, NOW())
	`, site.Id, result.Up)

	log.Info("Checking site ID", site.Id)
	return err
}

// Check checks a single site.
func Check(ctx context.Context, siteID int, svc *SiteService, db *sqlx.DB) error {
	site, err := svc.Site.Get(ctx, siteID)
	if err != nil {
		log.Error("Error checking for site ID", siteID, ":", err)
		return err
	}
	log.Info("Checking site ID", site.Id)

	return check(ctx, site, db)
}

func CheckAll(ctx context.Context, svc *SiteService, db *sqlx.DB) error {
	// Get all the tracked sites
	resp, err := svc.Site.List(ctx)
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
			log.Error("Error checking sites:", err)
			return check(ctx, site, db)
		})
	}
	return g.Wait()
}

// TODO: add a cron job that will use the CheckAll function every hour or so
