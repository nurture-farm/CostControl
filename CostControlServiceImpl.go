package CostControlSDK

import (
	"errors"
	"github.com/nurture-farm/CostControl/cache"
	"github.com/nurture-farm/CostControl/grafana"
	"github.com/nurture-farm/CostControl/models"
	ps "github.com/nurture-farm/CostControl/prometheus"
	"github.com/nurture-farm/CostControl/util"
)

func InitExpense(expense models.Expense) (bool, error) {
	if !util.IsConfigSet {
		return false, errors.New("library not yet initialised with config")
	}

	firingAlerts := cache.AlertsCache
	isAllowed := true

	if firingAlerts != nil && len(firingAlerts) > 0 {
		for _, alert := range firingAlerts {
			isAllowed = isAllowed && !isAlertForExpense(alert, expense)
		}
	}

	if !isAllowed {
		return isAllowed, errors.New("expense exceeds threshold")
	}

	err := ps.RecordExpense(expense)
	if err != nil {
		return false, err
	}

	return true, nil
}

func isAlertForExpense(alertParamMap map[string]string, expense models.Expense) bool {
	isAlert := true
	for key, value := range alertParamMap {
		isAlert = isAlert && expense.Tags[key] == value
	}
	return isAlert
}

func ConfigureThreshold(grafanaAlertRule models.GrafanaAlertRule) error {
	if !util.IsConfigSet {
		return errors.New("library not yet initialised with config")
	}

	return grafana.CreateAlert(grafanaAlertRule)
}

func Configure(config models.Config) {
	util.Config = config
	util.IsConfigSet = true
}
