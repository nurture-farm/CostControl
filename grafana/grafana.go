package grafana

import (
	"bytes"
	"encoding/json"
	"github.com/nurture-farm/costcontrol/httpclient"
	"github.com/nurture-farm/costcontrol/models"
	"github.com/nurture-farm/costcontrol/util"
	"io/ioutil"
	"net/http"
)

func CreateAlert(grafanaAlertRule models.GrafanaAlertRule) error {
	jsonData, err := json.Marshal(grafanaAlertRule)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPost, util.Config.GrafanaHost+util.GRAFANA_CREATE_ALERT+util.Config.GrafanaRulesDirectory, bytes.NewReader(jsonData))
	request.Header = http.Header{
		"Authorization": {"Bearer " + util.Config.ApiKey},
	}
	_, err = httpclient.Client.Do(request)
	return err
}

func GetFiringAlerts() ([]models.GrafanaFiringAlert, error) {

	client := httpclient.Client
	req, err := http.NewRequest("GET", util.Config.GrafanaHost+util.GRAFANA_GET_FIRING_ALERTS, nil)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"Authorization": {"Bearer " + util.Config.ApiKey},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result []models.GrafanaFiringAlert

	err = json.Unmarshal(body, &result)

	return result, err
}

func GetAlertRules() (map[string][]models.GrafanaAlertRule, error) {
	client := httpclient.Client
	req, err := http.NewRequest("GET", util.Config.GrafanaHost+util.GRAFANA_GET_RULES+util.Config.GrafanaRulesDirectory, nil)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"Authorization": {"Bearer " + util.Config.ApiKey},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result map[string][]models.GrafanaAlertRule
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		return nil, err
	}

	return result, nil
}
