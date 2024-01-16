package check

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kirre02/monitor-backend/internal/site/service"
	"golang.org/x/sync/errgroup"
)

type SiteService struct {
	Site service.SiteServiceInterface
}

func check(ctx context.Context, siteInfo *service.Site, db *sqlx.DB) error {
	// Perform a ping check on the site
	result, err := Ping(ctx, siteInfo.Url)
	if err != nil {
		return err
	}

	// Insert the result into the database
	_, err = db.ExecContext(ctx, `
		INSERT INTO checks (site_id, up, checked_at)
		VALUES ($1, $2, NOW())
	`, siteInfo.Id, result.Up)

	return err
}

// Check checks a single site.
func Check(ctx context.Context, siteID int, svc *SiteService, db *sqlx.DB) error {
	siteInfo, err := svc.Site.Get(ctx, siteID)
	if err != nil {
		return err
	}

	return check(ctx, siteInfo, db)
}

func CheckAll(ctx context.Context, svc *SiteService, db *sqlx.DB) error {
	// Get all the tracked sites
	resp, err := svc.Site.List(ctx)
	if err != nil {
		return err
	}

	// Check up to 8 sites concurrently.
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(8)
	for _, site := range resp.Sites {
		site := site
		g.Go(func() error {
			return check(ctx, site, db)
		})
	}
	return g.Wait()
}
