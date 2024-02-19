package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/kirre02/monitor-backend/internal/site/handler"
	"github.com/kirre02/monitor-backend/internal/site/service"
	"github.com/stretchr/testify/assert"
)

type Errs struct {
	Message []string `json:"message"`
}

func TestAddSiteHandler(t *testing.T) {

	type args struct {
		site *service.AddParams
	}

	type want struct {
		site   *service.Site
		status int
		err    error
		errs   Errs
	}

	type test struct {
		name string
		args
		want
	}

	tests := []test{
		{
			name: "normal",
			args: args{
				site: &service.AddParams{
					Url:  "example.com",
					Name: "example",
				},
			},
			want: want{
				site: &service.Site{
					Id:   1,
					Name: "example",
					Url:  "example.com",
				},
				err:    nil,
				status: http.StatusCreated,
			},
		},
		{
			name: "empty string",
			args: args{
				site: &service.AddParams{
					Url:  "example.com",
					Name: "",
				},
			},
			want: want{
				site: &service.Site{
					Id:   1,
					Name: "",
					Url:  "example.com",
				},
				err:    nil,
				status: http.StatusCreated,
			},
		},
		{
			name: "other error",
			args: args{
				site: &service.AddParams{},
			},
			want: want{
				site:   &service.Site{},
				err:    errors.New("Failed to add site\n"),
				status: http.StatusBadRequest,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			var err error

			if test.args.site != nil {
				err = json.NewEncoder(&buf).Encode(test.args.site)
			}

			if err != nil {
				t.Fatalf("Failed to encode test data: %s", err)
			}

			r := httptest.NewRequest(http.MethodPost, "/api/v1/site", &buf)
			w := httptest.NewRecorder()

			router := chi.NewRouter()

			svc := &service.SiteMock{
				CreateFunc: func(ctx context.Context, s *service.AddParams) (*service.Site, error) {
					return test.want.site, test.want.err
				},
			}

			siteHandler := &handler.SiteHandler{Svc: svc}
			router.Post("/api/v1/site", siteHandler.AddSite)

			router.ServeHTTP(w, r)
			resp := w.Result()

			// Check the response status code
			if resp.StatusCode != test.want.status {
				t.Errorf("Expected status code %d, got %d", test.want.status, resp.StatusCode)
			}

			// If the response status code is OK (200 or 201), decode the response body
			if resp.StatusCode != test.want.status {
				t.Errorf("Expected status code %d, got %d", test.want.status, resp.StatusCode)
			} else if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
				b, err := io.ReadAll(w.Body)
				assert.Nil(t, err)

				if len(test.want.errs.Message) > 0 {
					var errStruct Errs
					err = json.Unmarshal(b, &errStruct)
					assert.Nil(t, err)

					for i := range errStruct.Message {
						assert.Equal(t, test.want.errs.Message[i], errStruct.Message[i])
					}
				} else {
					errMsg := string(b)
					assert.Equal(t, test.want.err.Error(), errMsg)
				}
			}
		})
	}
}

func TestGetSiteHandler(t *testing.T) {
	type want struct {
		site   *service.Site
		status int
		err    error
	}

	type test struct {
		name     string
		urlParam string
		want
	}

	tests := []test{
		{
			name:     "Normal",
			urlParam: "1", // Assuming site ID 1 exists
			want: want{
				site: &service.Site{
					Id:   1,
					Name: "example",
					Url:  "example.com",
				},
				status: http.StatusOK,
				err:    nil,
			},
		},
		{
			name:     "Invalid ID",
			urlParam: "invalid", // An invalid ID
			want: want{
				site:   nil,
				status: http.StatusBadRequest,
				err:    nil,
			},
		},
		{
			name:     "Non-existing Site",
			urlParam: "100", // Assuming site ID 100 does not exist
			want: want{
				site:   nil,
				status: http.StatusInternalServerError,
				err:    errors.New("failed to retrieve site"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/site/%s", test.urlParam), nil)
			w := httptest.NewRecorder()

			router := chi.NewRouter()

			svc := &service.SiteMock{
				GetFunc: func(ctx context.Context, siteID int) (*service.Site, error) {
					return test.want.site, test.want.err
				},
			}

			siteHandler := &handler.SiteHandler{Svc: svc}
			router.Get("/api/v1/site/{id}", siteHandler.GetSite)

			router.ServeHTTP(w, r)
			resp := w.Result()

			if resp.StatusCode != test.want.status {
				t.Errorf("Expected status code %d, got %d", test.want.status, resp.StatusCode)
			}
		})
	}
}

func TestListSitesHandler(t *testing.T) {
	type want struct {
		sites  *service.ListResponse
		status int
		err    error
	}

	tests := []struct {
		name string
		want
	}{
		{
			name: "Normal",
			want: want{
				sites: &service.ListResponse{
					Sites: []*service.Site{
						{
							Id:   1,
							Name: "example1",
							Url:  "example1.com",
						},
						{
							Id:   2,
							Name: "example2",
							Url:  "example2.com",
						},
					},
				},
				status: http.StatusOK,
				err:    nil,
			},
		},
		{
			name: "Failed to Retrieve Sites",
			want: want{
				sites:  nil,
				status: http.StatusInternalServerError,
				err:    errors.New("failed to retrieve sites"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/api/v1/sites", nil)
			w := httptest.NewRecorder()

			router := chi.NewRouter()

			svc := &service.SiteMock{
				ListFunc: func(ctx context.Context) (*service.ListResponse, error) {
					return test.want.sites, test.want.err
				},
			}

			siteHandler := &handler.SiteHandler{Svc: svc}
			router.Get("/api/v1/sites", siteHandler.ListSites)

			router.ServeHTTP(w, r)
			resp := w.Result()

			if resp.StatusCode != test.want.status {
				t.Errorf("Expected status code %d, got %d", test.want.status, resp.StatusCode)
			}
		})
	}
}

func TestDeleteSiteHandler(t *testing.T) {
	type want struct {
		status int
		err    error
	}

	tests := []struct {
		name     string
		siteID   string
		want     want
		isNumber bool
	}{
		{
			name:     "Valid Site ID",
			siteID:   "1",
			want:     want{status: http.StatusOK, err: nil},
			isNumber: true,
		},
		{
			name:     "Invalid Site ID",
			siteID:   "invalid",
			want:     want{status: http.StatusBadRequest, err: nil},
			isNumber: false,
		},
		{
			name:     "Failed to Delete Site",
			siteID:   "2",
			want:     want{status: http.StatusInternalServerError, err: errors.New("failed to delete site")},
			isNumber: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/site/%s", test.siteID), nil)
			w := httptest.NewRecorder()

			router := chi.NewRouter()

			svc := &service.SiteMock{
				DeleteFunc: func(ctx context.Context, id int) error {
					if !test.isNumber {
						return fmt.Errorf("invalid site ID")
					}
					return test.want.err
				},
			}

			siteHandler := &handler.SiteHandler{Svc: svc}
			router.Delete("/api/v1/site/{id}", siteHandler.DeleteSite)

			router.ServeHTTP(w, r)
			resp := w.Result()

			if resp.StatusCode != test.want.status {
				t.Errorf("Expected status code %d, got %d", test.want.status, resp.StatusCode)
			}
		})
	}
}
