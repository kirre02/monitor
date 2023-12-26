package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/kirre02/monitor-backend/internal/site/service"
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
                err: nil,
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
