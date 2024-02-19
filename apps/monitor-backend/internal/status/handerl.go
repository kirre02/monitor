package status

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kirre02/monitor-backend/util"
)

type StatusHandler struct {
	Svc Service
}

func (sh *StatusHandler) StatusHandler(w http.ResponseWriter, r *http.Request) {
	statusResponse, err := sh.Svc.Status(r.Context())

	if err != nil {
		http.Error(w, "failed to check all sites", http.StatusInternalServerError)
		return
	}

	util.ToJson(w, statusResponse, http.StatusOK)
}

func (sh *StatusHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", sh.StatusHandler)

	return r
}
