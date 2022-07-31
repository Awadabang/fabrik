package registry

type RegistryReq struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}

type DestroyReq struct {
	Name string `json:"name"`
}
