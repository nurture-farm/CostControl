package models

type GrafanaFiringAlert struct {
	Annotations  GrafanaAnnotations `json:"annotations"`
	EndsAt       string             `json:"endsAt"`
	Fingerprint  string             `json:"fingerprint"`
	Receivers    []GrafanaReceiver  `json:"receivers"`
	StartsAt     string             `json:"startsAt"`
	Status       GrafanaAlertStatus `json:"status"`
	UpdatedAt    string             `json:"updatedAt"`
	GeneratorURL string             `json:"generatorURL"`
	Labels       map[string]string  `json:"labels"`
}

type GrafanaReceiver struct {
	Name string `json:"name"`
}

type GrafanaAlertStatus struct {
	InhibitedBy []string `json:"inhibited_by"`
	SilencedBy  []string `json:"silenced_by"`
	State       string   `json:"state"`
}

type GrafanaAnnotations struct {
	ValueString string `json:"___value_string___"`
}
