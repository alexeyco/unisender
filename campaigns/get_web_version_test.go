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

func TestGetWebVersionRequest_Execute(t *testing.T) {
	expectedCampaignID := test.RandomInt64(9999, 999999)
	var givenCampaignID int64

	expectedResult := &campaigns.GetWebVersionResult{
		LetterID:      test.RandomInt64(9999, 999999),
		WebLetterLink: test.RandomString(12, 36),
	}

	var givenResult *campaigns.GetWebVersionResult

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

	givenResult, err := campaigns.GetWebVersion(req, expectedCampaignID).
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
