package check

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kirre02/monitor-backend/util"
)

type CheckHandler struct {
	Svc *Service
}

// CheckHandler handles HTTP requests to check a single site.
func (ch *CheckHandler) CheckHandler(w http.ResponseWriter, r *http.Request) {
	siteIDStr := chi.URLParam(r, "id")
	siteID, err := strconv.Atoi(siteIDStr)
	if err != nil {
		http.Error(w, "invalid Site ID", http.StatusBadRequest)
		return
	}

	pingResponse, err := ch.Svc.Check(r.Context(), siteID)
	if err != nil {
		http.Error(w, "failed to check site", http.StatusInternalServerError)
		return
	}

	util.ToJson(w, pingResponse, http.StatusOK)
}

// CheckAllHandler handles HTTP requests to check all sites.
func (ch *CheckHandler) CheckAllHandler(w http.ResponseWriter, r *http.Request) {
	err := ch.Svc.CheckAll(r.Context())
	if err != nil {
		http.Error(w, "failed to check all sites", http.StatusInternalServerError)
		return
	}

	// Assuming you want to return a simple success JSON response
	response := map[string]string{"status": "success"}

	util.ToJson(w, response, http.StatusOK)
}

func (ch *CheckHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/{id}", ch.CheckHandler)
	r.Get("/all", ch.CheckAllHandler)

	return r
}
