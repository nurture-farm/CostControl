package prometheus

import (
	"github.com/nurture-farm/costcontrol/models"
	"github.com/nurture-farm/costcontrol/util"
	"github.com/prometheus/client_golang/prometheus"
	v "gopkg.in/go-playground/validator.v9"
	m "k8s.io/utils/keymutex"
)

var metrics = make(map[string]*prometheus.CounterVec)
var metricTagsMap = make(map[string][]string)

var validator = v.New()
var keyMutex = m.NewHashed(util.MAX_LOCKS)

func RecordExpense(expense models.Expense) error {
	err := validator.Struct(expense)
	if err != nil {
		return err
	}

	metric := metrics[expense.ExpenseName]

	if metric == nil {
		keyMutex.LockKey(expense.ExpenseName)
		metric = metrics[expense.ExpenseName]
		if metric == nil {
			metric = registerMetric(expense)
			metrics[expense.ExpenseName] = metric
		}
		err := keyMutex.UnlockKey(expense.ExpenseName)
		if err != nil {
			return err
		}
	}

	metricTagList := metricTagsMap[expense.ExpenseName]

	tagValues := getTagsInOrder(expense.Tags, metricTagList)

	metric.WithLabelValues(tagValues...).Add(float64(expense.Value))

	return nil
}

func registerMetric(expense models.Expense) *prometheus.CounterVec {
	tags := make([]string, len(expense.Tags))
	index := 0
	for metricName, _ := range expense.Tags {
		tags[index] = metricName
		index++
	}

	metrics[expense.ExpenseName] = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_Budgets" + expense.ExpenseName,
		Help: "Summary metrics for Budgets",
	}, tags)
	metricTagsMap[expense.ExpenseName] = tags
	prometheus.MustRegister(metrics[expense.ExpenseName])

	return metrics[expense.ExpenseName]
}

func getTagsInOrder(expenseTagValsMap map[string]string, metricTagsList []string) []string {
	tagVals := make([]string, len(metricTagsList))

	for index, tag := range metricTagsList {
		tagVals[index] = expenseTagValsMap[tag]
	}

	return tagVals
}
