package service_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/kirre02/monitor-backend/internal/site/service"
	"github.com/stretchr/testify/assert"
)

func TestSiteService_Create(t *testing.T) {
	type args struct {
		site *service.AddParams
	}

	type want struct {
		site             *service.Site
		err              error
		recordNotCreated bool
	}

	type test struct {
		name       string
		beforeTest func(sqlmock.Sqlmock)
		args
		want
	}

	tests := []test{
		{
			name: "",
			args: args{nil},
			want: want{
				site: nil,
				err:  fmt.Errorf("request cannot be nil"),
			},
		},
		{
			name: "empty site name",
			args: args{
				site: &service.AddParams{
					Name: "",
					Url:  "example.com",
				},
			},
			want: want{
				site: &service.Site{
					Id:   1,
					Name: "",
					Url:  "example.com",
				},
				err:              nil,
				recordNotCreated: true,
			},
		},
		{
			name: "Normal",
			args: args{
				site: &service.AddParams{
					Url:  "example.com",
					Name: "example",
				},
			},
			want: want{
				site: &service.Site{
					Id:   2,
					Name: "example",
					Url:  "example.com",
				},
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			s := &service.Service{DB: db}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			// Set up an expectation for the NamedExecContext call
			if tt.args.site != nil && !tt.recordNotCreated {
				mockSQL.ExpectExec("INSERT INTO sites (.+) VALUES (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
			}

			created, err := s.Add(ctx, tt.args.site)

			if created == nil {
				return
			}

			if (err != nil) != tt.recordNotCreated {
				t.Errorf("error = %v, wantErr %v", err, tt.recordNotCreated)
				return
			}
		})
	}
}

func TestSiteService_Get(t *testing.T) {
	type args struct {
		siteID int
	}

	type want struct {
		site *service.Site
		err  error
	}

	type test struct {
		name       string
		beforeTest func(sqlmock.Sqlmock)
		args
		want
	}

	tests := []test{
		{
			name: "Get site by ID",
			args: args{
				siteID: 1,
			},
			want: want{
				site: &service.Site{
					Id:   1,
					Name: "example",
					Url:  "example.com",
				},
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			s := &service.Service{DB: db}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			rows := sqlmock.NewRows([]string{"id", "url", "name"}).
				AddRow(tt.want.site.Id, tt.want.site.Url, tt.want.site.Name)

			mockSQL.ExpectQuery("SELECT id, url, name FROM sites WHERE id = ?").
				WithArgs(tt.args.siteID).
				WillReturnRows(rows)

			got, err := s.Get(ctx, tt.args.siteID)

			assert.Equal(t, tt.want.err, err)

			assert.Equal(t, tt.want.site.Id, got.Id)
			assert.Equal(t, tt.want.site.Url, got.Url)
			assert.Equal(t, tt.want.site.Name, got.Name)
		})
	}
}
func TestSiteService_List(t *testing.T) {
	type want struct {
		sites []*service.Site
		err   error
	}

	type test struct {
		name       string
		beforeTest func(sqlmock.Sqlmock)
		want
	}

	tests := []test{
		{
			name: "Get list of sites",
			want: want{
				sites: []*service.Site{
					{
						Id:   1,
						Url:  "example1.com",
						Name: "example1",
					},
					{
						Id:   2,
						Url:  "example2.com",
						Name: "example2",
					},
				},
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			s := &service.Service{DB: db}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			rows := sqlmock.NewRows([]string{"id", "url", "name"})
			for _, site := range tt.want.sites {
				rows.AddRow(site.Id, site.Url, site.Name)
			}

			mockSQL.ExpectQuery("SELECT id, Url, Name FROM sites").
				WillReturnRows(rows)

			got, err := s.List(ctx)

			assert.Equal(t, tt.want.err, err)
			assert.Equal(t, tt.want.sites, got.Sites)
		})
	}
}

func TestSiteService_Delete(t *testing.T) {
	type args struct {
		ctx    context.Context
		SiteID int
	}

	type test struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		wantErr    bool
	}

	tests := []test{
		{
			name: "Delete site with ID 1",
			args: args{
				ctx:    context.Background(),
				SiteID: 1,
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectExec("(?i)DELETE FROM sites WHERE id = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Delete non-existing site",
			args: args{
				ctx:    context.Background(),
				SiteID: 5,
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectExec("(?i)DELETE FROM sites WHERE id = ?").WithArgs(5).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")
			s := &service.Service{DB: db}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			err := s.Delete(tt.args.ctx, tt.args.SiteID)

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
