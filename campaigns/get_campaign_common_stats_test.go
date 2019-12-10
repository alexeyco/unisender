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

func TestGetCampaignCommonStatsRequest_Execute(t *testing.T) {
	expectedCampaignID := test.RandomInt64(9999, 999999)
	var givenCampaignID int64

	expectedResult := &campaigns.GetCampaignStatsResult{
		Total:         test.RandomInt64(9999, 999999),
		Sent:          test.RandomInt64(9999, 999999),
		Delivered:     test.RandomInt64(9999, 999999),
		ReadUnique:    test.RandomInt64(9999, 999999),
		ReadAll:       test.RandomInt64(9999, 999999),
		ClickedUnique: test.RandomInt64(9999, 999999),
		ClickedAll:    test.RandomInt64(9999, 999999),
		Unsubscribed:  test.RandomInt64(9999, 999999),
		Spam:          test.RandomInt64(9999, 999999),
	}

	var givenResult *campaigns.GetCampaignStatsResult

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

	err := campaigns.CancelCampaign(req, expectedCampaignID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedCampaignID != givenCampaignID {
		t.Fatalf(`Campaign ID should be %d, %d given`, expectedCampaignID, givenCampaignID)
	}

	if reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatalf("Results should be equal")
	}
}
