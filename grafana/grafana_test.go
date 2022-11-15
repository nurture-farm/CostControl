package grafana

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/nurture-farm/costcontrol/httpclient"
	"github.com/nurture-farm/costcontrol/mocks"
	"github.com/nurture-farm/costcontrol/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func init() {
	httpclient.Client = &mocks.MockClient{}
}

func TestCreateAlert(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, nil
	}

	err := CreateAlert(models.GrafanaAlertRule{
		Name:     "cc_campaign",
		Interval: "5m",
		Rules:    nil,
	})

	if err != nil {
		t.Errorf("create alert not working as expected")
	}
}

func TestCreateAlertHttpError(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New("http error")
	}

	err := CreateAlert(models.GrafanaAlertRule{
		Name:     "cc_campaign",
		Interval: "5m",
		Rules:    nil,
	})

	if err == nil {
		t.Errorf("create alert not working as expected for negative case")
	}
}

func TestGetFiringAlerts(t *testing.T) {
	response := [1]models.GrafanaFiringAlert{
		{
			Annotations:  models.GrafanaAnnotations{},
			EndsAt:       "",
			Fingerprint:  "",
			Receivers:    nil,
			StartsAt:     "",
			Status:       models.GrafanaAlertStatus{},
			UpdatedAt:    "",
			GeneratorURL: "",
			Labels:       nil,
		},
	}
	data, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	r := ioutil.NopCloser(bytes.NewReader(data))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	resp, err := GetFiringAlerts()

	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestGetFiringAlertsError(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New("http error")
	}

	resp, err := GetFiringAlerts()

	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestGetFiringAlertsInvalidResponse(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"invalid": "json"}`)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	resp, err := GetFiringAlerts()

	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestGetAlertRules(t *testing.T) {
	response := map[string][]models.GrafanaAlertRule{
		"Platform": {
			{
				Name:     "",
				Interval: "",
			}},
	}
	data, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	r := ioutil.NopCloser(bytes.NewReader(data))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	resp, err := GetAlertRules()

	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestGetAlertRulesError(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New("http error")
	}

	resp, err := GetAlertRules()

	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestGetAlertRulesInvalidResponse(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"invalid": "json"}`)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	resp, err := GetAlertRules()

	assert.NotNil(t, err)
	assert.Nil(t, resp)
}
