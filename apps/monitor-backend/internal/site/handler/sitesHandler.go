package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kirre02/monitor-backend/internal/site/service"
)

type SiteHandler struct {
	Service service.SiteServiceInterface
}

func (sh *SiteHandler) AddSite(w http.ResponseWriter, r *http.Request) {
	var params service.AddParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	addSite, err := sh.Service.Add(r.Context(), &params)
	if err != nil {
		http.Error(w, "Failed to add site", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(addSite)
}

func (sh *SiteHandler) GetSite(w http.ResponseWriter, r *http.Request) {
	siteIDStr := chi.URLParam(r, "id")
	siteID, err := strconv.Atoi(siteIDStr)
	if err != nil {
		http.Error(w, "invalid Site ID", http.StatusBadRequest)
	}

	getSite, err := sh.Service.Get(r.Context(), siteID)
	if err != nil {
		http.Error(w, "failed to retrieve site", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getSite)
}

func (sh *SiteHandler) DeleteSite(w http.ResponseWriter, r *http.Request) {
	siteIDStr := chi.URLParam(r, "id")
	siteID, err := strconv.Atoi(siteIDStr)
	if err != nil {
		http.Error(w, "invalid Site ID", http.StatusBadRequest)
	}

	err = sh.Service.Delete(r.Context(), siteID)
	if err != nil {
		http.Error(w, "failed to retrieve site", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Site was successfully deleted"})
}

func (sh *SiteHandler) ListSites(w http.ResponseWriter, r *http.Request) {
	sites, err := sh.Service.List(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve sites", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sites)
}

func (sh *SiteHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/sites/{id}", sh.GetSite)
	r.Get("/sites", sh.ListSites)
	r.Post("/site", sh.AddSite)
	r.Delete("/sites/{id}", sh.DeleteSite)

	return r
}
