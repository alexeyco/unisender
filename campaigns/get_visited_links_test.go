package campaigns_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/campaigns"
	"github.com/alexeyco/unisender/test"
)

func TestGetVisitedLinksRequest_Group(t *testing.T) {
	expectedCampaignID := test.RandomInt64(9999, 999999)
	expectedResult := randomGetVisitedLinksResult()

	expectedGroup := 1
	var givenGroup int

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenGroup, _ = strconv.Atoi(req.FormValue("group"))

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.GetVisitedLinks(req, expectedCampaignID).
		Group().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedGroup != givenGroup {
		t.Fatalf(`Group should be %d, %d given`, expectedGroup, givenGroup)
	}
}

func TestGetVisitedLinksRequest_Execute(t *testing.T) {
	expectedCampaignID := test.RandomInt64(9999, 999999)
	var givenCampaignID int64

	expectedResult := randomGetVisitedLinksResult()
	var givenResult *campaigns.GetVisitedLinksResult

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenCampaignID, _ = strconv.ParseInt(req.FormValue("campaign_id"), 10, 64)

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := campaigns.GetVisitedLinks(req, expectedCampaignID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedCampaignID != givenCampaignID {
		t.Fatalf(`Campaign ID should be %d, %d given`, expectedCampaignID, givenCampaignID)
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal("Results should be equal")
	}
}

func randomGetVisitedLinksResult() *campaigns.GetVisitedLinksResult {
	f := test.RandomInt(4, 8)
	d := test.RandomInt(12, 32)

	fields := make([]string, f)
	for i := 0; i < f; i++ {
		fields[i] = test.RandomString(12, 36)
	}

	data := make([][]string, d)
	for i := 0; i < d; i++ {
		row := make([]string, f)
		for j := 0; j < f; j++ {
			row[j] = test.RandomString(12, 36)
		}

		data[i] = row
	}

	return &campaigns.GetVisitedLinksResult{
		Fields: fields,
		Data:   data,
	}
}
