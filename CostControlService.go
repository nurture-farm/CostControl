package CostControlSDK

import "github.com/nurture-farm/costcontrol/models"

type CostControlService interface {
	Configure(config models.Config)
	InitExpense(expense models.Expense) (bool, error)
	ConfigureThreshold(grafanaAlertRule models.GrafanaAlertRule) error
}
