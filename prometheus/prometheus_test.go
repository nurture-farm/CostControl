package prometheus

import (
	"github.com/nurture-farm/costcontrol/models"
	"testing"
)

func TestRecordExpenseInvalidInput(t *testing.T) {
	err := RecordExpense(models.Expense{
		ExpenseName: "",
		Value:       -10,
		Tags:        nil,
	})

	if err == nil {
		t.Errorf("validation of expense not working as expected")
	}
}

func TestRecordExpenseNewExpenseName(t *testing.T) {
	err := RecordExpense(models.Expense{
		ExpenseName: "Campaign",
		Value:       100,
		Tags:        nil,
	})

	if err != nil {
		t.Errorf("record expense for new expense name not working as expected")
	}
}

func TestRecordExpenseExistingExpenseName(t *testing.T) {
	err := RecordExpense(models.Expense{
		ExpenseName: "Campaign",
		Value:       100,
		Tags: map[string]string{"cost_center": "marketing",
			"service":      "campaign",
			"comm_channel": "APP_NOTIFICATION",
			"campaign_id":  "102"},
	})

	err = RecordExpense(models.Expense{
		ExpenseName: "Campaign",
		Value:       120,
		Tags: map[string]string{"cost_center": "marketing",
			"service":      "campaign",
			"comm_channel": "APP_NOTIFICATION",
			"campaign_id":  "102"},
	})

	if err != nil {
		t.Errorf("record expense for existing expense name not working as expected")
	}
}
