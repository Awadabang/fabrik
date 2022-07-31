package registry

type RegistryReq struct {
	Name      string `json:"name"`
	Addr      string `json:"addr"`
	AccessKey string `json:"accessKey"`
}

type DestroyReq struct {
	Name string `json:"name"`
}
