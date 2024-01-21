package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kirre02/monitor-backend/internal/site/service"
	"github.com/kirre02/monitor-backend/util"
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

	util.ToJson(w, addSite, http.StatusCreated)
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

	util.ToJson(w, getSite, http.StatusOK)
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

	resp := map[string]string{"message": "Site was successfully deleted"}
	util.ToJson(w, resp, http.StatusOK)
}

func (sh *SiteHandler) ListSites(w http.ResponseWriter, r *http.Request) {
	sites, err := sh.Service.List(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve sites", http.StatusInternalServerError)
		return
	}

	util.ToJson(w, sites, http.StatusOK)
}

func (sh *SiteHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/get/{id}", sh.GetSite)
	r.Get("/list", sh.ListSites)
	r.Post("/add", sh.AddSite)
	r.Delete("/del/{id}", sh.DeleteSite)

	return r
}
