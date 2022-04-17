package version

type HealthCheck struct {
	App     string `json:"app"`
	Version string `json:"version"`
	Env     string `json:"env"`
	Author  string `json:"author"`
}
