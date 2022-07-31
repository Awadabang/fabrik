package handlers

import (
	"io"
	"net/http"

	"github.com/Awadabang/fabrik/internal/services"
)

func LogHandler(svcCtx *services.Svc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			svcCtx.LogService.Write(string(msg))
		default:
		}
	}
}
