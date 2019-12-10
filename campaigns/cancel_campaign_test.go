package campaigns

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/test"
)

func TestCancelCampaignRequest_Execute(t *testing.T) {
	expectedCampaignID := test.RandomInt64(9999, 999999)
	var givenCampaignID int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenCampaignID, _ = strconv.ParseInt(req.FormValue("campaign_id"), 10, 64)

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := CancelCampaign(req, expectedCampaignID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedCampaignID != givenCampaignID {
		t.Fatalf(`Campaign ID should be %d, %d given`, expectedCampaignID, givenCampaignID)
	}
}
