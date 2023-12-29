package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
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

			siteHandler := &handler.SiteHandler{Service: svc}
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
