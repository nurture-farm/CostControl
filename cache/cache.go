package cache

import (
	"errors"
	"github.com/nurture-farm/CostControl/grafana"
	"github.com/nurture-farm/CostControl/models"
	"github.com/nurture-farm/CostControl/util"
	"strings"
	"time"
)

var AlertsCache []map[string]string

func init() {
	err := populateCache()
	if err != nil {
		panic("unable to initialize alerts cache")
	}
	if util.Config.AlertStatusRefreshStrategy == models.TTL {
		triggerPopulateCache()
	}
}

func triggerPopulateCache() {
	ticker := time.NewTicker(util.Config.AlertStatusTTL)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				populateCache()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func populateCache() error {
	firingAlerts, err := grafana.GetFiringAlerts()
	if err != nil {
		return err
	}

	alertConfigs, err := GetAlertConfigs()
	if err != nil {
		return err
	}

	alerts := make([]map[string]string, 0, len(firingAlerts))
	index := 0
	for _, firingAlert := range firingAlerts {
		alertName := firingAlert.Labels[util.ALERT_NAME_LABEL]
		if !strings.Contains(alertName, util.GRAFANA_ALERT_RULE_NAME_PREFIX) {
			continue
		}
		alertRuleExpr := alertConfigs[alertName]
		pairsForAlert, err := getPairs(alertRuleExpr)
		if err != nil {
			continue
		}
		alerts = alerts[0 : index+1]
		alerts[index] = pairsForAlert
		index++
	}
	AlertsCache = alerts
	return err
}

func GetAlertConfigs() (map[string]string, error) {
	var res = make(map[string]string)

	alertRules, err := grafana.GetAlertRules()
	if err != nil {
		return nil, err
	}

	platformAlertRules := alertRules[util.Config.GrafanaRulesDirectory]

	for _, alertRule := range platformAlertRules {
		name := alertRule.Name
		expr := alertRule.Rules[util.GRAFANA_EXPRESSION_RULE_INDEX].GrafanaAlert.Data[util.GRAFANA_EXPRESSION_RULE_DATA_INDEX].Model.Expr
		res[name] = expr
	}

	return res, nil
}

func getPairs(expression string) (map[string]string, error) {
	expression = strings.Split(expression, "}")[0]
	expression = strings.Split(expression, "{")[1]

	pairs := strings.Split(expression, util.GRAFANA_EXPRESSION_SEPARATOR)

	var pairMap = make(map[string]string)
	for _, pair := range pairs {
		if !strings.Contains(pair, util.GRAFANA_EXPRESSSION_COMPARISON_OP) {
			return nil, errors.New("unrecognized alert rule")
		}
		key := strings.Split(pair, util.GRAFANA_EXPRESSSION_COMPARISON_OP)[util.GRAFANA_EXPRESSION_KEY_INDEX]
		val := strings.Split(pair, util.GRAFANA_EXPRESSSION_COMPARISON_OP)[util.GRAFANA_EXPRESSION_VALUE_INDEX]
		pairMap[key] = unquote(val)
	}

	return pairMap, nil
}

/*
Converts "Cost" to Cost
*/
func unquote(str string) string {
	return strings.Split(str, "\"")[1]
}
