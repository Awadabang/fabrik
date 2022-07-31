package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Awadabang/fabrik/internal/services"
	"github.com/Awadabang/fabrik/internal/services/registry"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("error when reading body"))
		return
	}
	w.Write([]byte("根目录"))
	w.Write(body)
}

func RegisterHandler(svcCtx *services.Svc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("error when reading body"))
			return
		}
		req := registry.RegistryReq{}
		err = json.Unmarshal(bodyBytes, &req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Write([]byte("注册请求为："))
		w.Write(bodyBytes)

		svcCtx.Registry.Add(req.Name, req.Addr)
	}
}
