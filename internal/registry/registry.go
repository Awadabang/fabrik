package registry

import (
	"encoding/json"
	"io"
	"net/http"
)

type Registry struct {
	registryServer *http.Server
}

type RegistryOption func(*Registry)

func NewRegistry(opts ...RegistryOption) *Registry {
	registry := &Registry{}
	for _, appply := range opts {
		appply(registry)
	}

	return registry
}

func (r *Registry) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/register", registerHandler)

	r.registryServer = &http.Server{
		Addr:    "0.0.0.0:8111",
		Handler: mux,
	}

	r.registryServer.ListenAndServe()
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("error when reading body"))
		return
	}
	w.Write([]byte("根目录"))
	w.Write(body)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("error when reading body"))
		return
	}
	req := RegistryReq{}
	err = json.Unmarshal(bodyBytes, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte("注册请求为："))
	w.Write(bodyBytes)
}
