package CostControlSDK

import "github.com/nurture-farm/CostControl/models"

type CostControlService interface {
	Configure(config models.Config)
	InitExpense(expense models.Expense) (bool, error)
	ConfigureThreshold(grafanaAlertRule models.GrafanaAlertRule) error
}
