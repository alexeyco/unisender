package campaigns_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/campaigns"
	"github.com/alexeyco/unisender/test"
)

func TestGetCampaignsRequest_From(t *testing.T) {
	expectedResult := randomGetCampaignsResult()

	expectedFrom := test.RandomTime(12, 365)
	var givenFrom time.Time

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenFrom, _ = time.Parse("2006-01-02 15:04:05", req.FormValue("from"))

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.GetCampaigns(req).
		From(expectedFrom).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedFrom != givenFrom {
		t.Fatalf(`From should be "%s", "%s" given`, expectedFrom, givenFrom)
	}
}

func TestGetCampaignsRequest_To(t *testing.T) {
	expectedResult := randomGetCampaignsResult()

	expectedTo := test.RandomTime(12, 365)
	var givenTo time.Time

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTo, _ = time.Parse("2006-01-02 15:04:05", req.FormValue("to"))

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.GetCampaigns(req).
		To(expectedTo).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTo != givenTo {
		t.Fatalf(`To should be "%s", "%s" given`, expectedTo, givenTo)
	}
}

func TestGetCampaignsRequest_Limit(t *testing.T) {
	expectedResult := randomGetCampaignsResult()

	expectedLimit := test.RandomInt(9999, 999999)
	var givenLimit int

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLimit, _ = strconv.Atoi(req.FormValue("limit"))

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.GetCampaigns(req).
		Limit(expectedLimit).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLimit != givenLimit {
		t.Fatalf(`Limit should be %d, %d given`, expectedLimit, givenLimit)
	}
}

func TestGetCampaignsRequest_Offset(t *testing.T) {
	expectedResult := randomGetCampaignsResult()

	expectedOffset := test.RandomInt(9999, 999999)
	var givenOffset int

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenOffset, _ = strconv.Atoi(req.FormValue("offset"))

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.GetCampaigns(req).
		Offset(expectedOffset).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedOffset != givenOffset {
		t.Fatalf(`Offset should be %d, %d given`, expectedOffset, givenOffset)
	}
}

func TestGetCampaignsRequest_Execute(t *testing.T) {
	expectedResult := randomGetCampaignsResult()
	var givenResult []campaigns.GetCampaignsResult

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := campaigns.GetCampaigns(req).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal("Results should be equal")
	}
}

func randomGetCampaignsResult() []campaigns.GetCampaignsResult {
	num := test.RandomInt(12, 36)
	res := make([]campaigns.GetCampaignsResult, num)

	for i := 0; i < num; i++ {
		res[i] = campaigns.GetCampaignsResult{
			ID:          test.RandomInt64(9999, 999999),
			StartTime:   test.RandomTime(12, 365),
			Status:      test.RandomString(12, 36),
			MessageID:   test.RandomInt64(9999, 999999),
			ListID:      test.RandomInt64(9999, 999999),
			Subject:     test.RandomString(12, 36),
			SenderName:  test.RandomString(12, 36),
			SenderEmail: test.RandomString(12, 36),
			StatsURL:    test.RandomString(12, 36),
		}
	}

	return res

}
