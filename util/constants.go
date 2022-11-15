package util

import "github.com/nurture-farm/CostControl/models"

const (
	GRAFANA_GET_RULES                  = "/api/ruler/grafana/api/v1/rules/"
	GRAFANA_GET_FIRING_ALERTS          = "/api/alertmanager/grafana/api/v2/alerts"
	GRAFANA_CREATE_ALERT               = "/api/ruler/grafana/api/v1/rules/"
	GRAFANA_EXPRESSION_RULE_INDEX      = 0
	GRAFANA_EXPRESSION_RULE_DATA_INDEX = 0
	GRAFANA_EXPRESSION_SEPARATOR       = ","
	GRAFANA_EXPRESSSION_COMPARISON_OP  = "="
	GRAFANA_EXPRESSION_KEY_INDEX       = 0
	GRAFANA_EXPRESSION_VALUE_INDEX     = 1
	GRAFANA_ALERT_RULE_NAME_PREFIX     = "cc_"
	ALERT_NAME_LABEL                   = "alertname"
	MAX_LOCKS                          = 20
)

var Config models.Config
var IsConfigSet bool
