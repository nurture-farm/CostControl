package models

type GrafanaAlertRule struct {
	Name     string      `json:"name"`
	Interval string      `json:"interval"`
	Rules    []AlertRule `json:"rules"`
}

type AlertRule struct {
	Expr         string       `json:"expr"`
	ForDuration  string       `json:"for"`
	GrafanaAlert GrafanaAlert `json:"grafana_alert"`
}

type GrafanaAlert struct {
	Id              int64  `json:"id"`
	OrgId           int64  `json:"orgId"`
	Title           string `json:"title"`
	Condition       string `json:"condition"`
	Data            []Data `json:"data"`
	Updated         string `json:"updated"`
	IntervalSeconds int64  `json:"intervalSeconds"`
	Version         int64  `json:"version"`
	Uid             string `json:"uid"`
	NamespaceUid    string `json:"namespace_uid"`
	NamespaceId     int64  `json:"namespace_id"`
	RuleGroup       string `json:"rule_group"`
	NoDataState     string `json:"no_data_state"`
	ExecErrState    string `json:"exec_err_state"`
}

type Data struct {
	RefId             string            `json:"refId"`
	QueryType         string            `json:"queryType"`
	RelativeTimeRange RelativeTimeRange `json:"relativeTimeRange"`
	DatasourceUid     string            `json:"datasourceUid"`
	Model             Model             `json:"model"`
}

type Model struct {
	Conditions    []Condition `json:"conditions"`
	Datasource    Datasource  `json:"datasource"`
	Downsampler   string      `json:"downsampler"`
	Hide          bool        `json:"hide"`
	IntervalMs    int64       `json:"intervalMs"`
	MaxDataPoints int64       `json:"maxDataPoints"`
	RefId         string      `json:"refId"`
	Type          string      `json:"type"`
	UpSampler     string      `json:"upSampler"`
	LegendFormat  string      `json:"legendFormat"`
	Interval      string      `json:"interval"`
	Expr          string      `json:"expr"`
	Exemplar      bool        `json:"exemplar"`
}

type Condition struct {
	Evaluator Evaluator `json:"evaluator"`
	Operator  Operator  `json:"operator"`
	Query     Query     `json:"query"`
	Reducer   Reducer   `json:"reducer"`
	Type      string    `json:"type"`
}

type Evaluator struct {
	Params []int64 `json:"params"`
	Type   string  `json:"type"`
}
type Operator struct {
	Type string `json:"type"`
}
type Query struct {
	Params []string `json:"params"`
}
type Reducer struct {
	Params []string `json:"params"`
	Type   string   `json:"type"`
}
type Datasource struct {
	Type string `json:"type"`
	Uid  string `json:"uid"`
}

type Duration struct {
	Duration int64
}

type RelativeTimeRange struct {
	from Duration
	to   Duration
}
