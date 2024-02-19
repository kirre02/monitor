package status

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	DB *sqlx.DB
}

// describes the current status of a site and when it was last checked
type SiteStatus struct {
	Up        bool      `json:"up"`
	CheckedAt time.Time `json:"checked_at"`
}

// StatusResponse describes the response from the status endpoint
type StatusResponse struct {
	// Sites containes the current status of all sites, keyed to the site ID
	Sites map[int]SiteStatus `json:"sites"`
}

func NewStatusService(db *sqlx.DB) *Service {
	return &Service{DB: db}
}

//Status checks the current status (up, down or unknown) of all sites

func (s *Service) Status(ctx context.Context) (*StatusResponse, error) {
	rows, err := s.DB.QueryContext(ctx, `
	SELECT DISTINCT ON (site_id) site_id, up, checked_at
	FROM checks
	ORDER BY site_id, checked_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]SiteStatus)
	for rows.Next() {
		var siteID int
		var status SiteStatus
		if err := rows.Scan(&siteID, &status.Up, &status.CheckedAt); err != nil {
			return nil, err
		}
		result[siteID] = status
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &StatusResponse{Sites: result}, nil
}
