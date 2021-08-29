package handlers

import (
	"net/http"

	log "github.com/nstoker/fictional-pancake/internal/app_logger"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Logger.Errorf("Failed request %+v", r)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
