package campaigns_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/campaigns"
	"github.com/alexeyco/unisender/test"
)

func TestCreateCampaignRequest_StartTime(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	expectedResult := randomCreateCampaignResult()

	expectedStartTime := test.RandomTime(12, 365).Round(time.Minute)
	var givenStartTime time.Time

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenStartTime, _ = time.Parse("2006-01-02 15:04", req.FormValue("start_time"))

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.CreateCampaign(req, expectedMessageID).
		StartTime(expectedStartTime).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedStartTime != givenStartTime {
		t.Fatalf(`Start time should be "%s", "%s" given`, expectedStartTime, givenStartTime)
	}
}

func TestCreateCampaignRequest_TrackRead(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	expectedResult := randomCreateCampaignResult()

	expectedTrackRead := 1
	var givenTrackRead int

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTrackRead, _ = strconv.Atoi(req.FormValue("track_read"))

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.CreateCampaign(req, expectedMessageID).
		TrackRead().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTrackRead != givenTrackRead {
		t.Fatalf(`Param "track_read" should be %d, %d given`, expectedTrackRead, givenTrackRead)
	}
}

func TestCreateCampaignRequest_TrackLinks(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	expectedResult := randomCreateCampaignResult()

	expectedTrackLinks := 1
	var givenTrackLinks int

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTrackLinks, _ = strconv.Atoi(req.FormValue("track_links"))

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.CreateCampaign(req, expectedMessageID).
		TrackLinks().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTrackLinks != givenTrackLinks {
		t.Fatalf(`Param "track_links" should be %d, %d given`, expectedTrackLinks, givenTrackLinks)
	}
}

func TestCreateCampaignRequest_Contacts(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	expectedResult := randomCreateCampaignResult()

	expectedContacts := test.RandomStringSlice(12, 36)
	var givenContacts []string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenContacts = strings.Split(req.FormValue("contacts"), ",")

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.CreateCampaign(req, expectedMessageID).
		Contacts(expectedContacts...).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedContacts, givenContacts) {
		t.Fatal("Contacts should be equal")
	}
}

func TestCreateCampaignRequest_ContactsURL(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	expectedResult := randomCreateCampaignResult()

	expectedContactsURL := test.RandomString(12, 36)
	var givenContactsURL string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenContactsURL = req.FormValue("contacts_url")

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.CreateCampaign(req, expectedMessageID).
		ContactsURL(expectedContactsURL).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedContactsURL != givenContactsURL {
		t.Fatalf(`Contacts URL should be "%s", "%s" given`, expectedContactsURL, givenContactsURL)
	}
}

func TestCreateCampaignRequest_TrackGoogleAnalytics(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	expectedResult := randomCreateCampaignResult()

	expectedTrackGoogleAnalytics := 1
	var givenTrackGoogleAnalytics int

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTrackGoogleAnalytics, _ = strconv.Atoi(req.FormValue("track_ga"))

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.CreateCampaign(req, expectedMessageID).
		TrackGoogleAnalytics().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTrackGoogleAnalytics != givenTrackGoogleAnalytics {
		t.Fatalf(`Param "track_ga" should be %d, %d given`, expectedTrackGoogleAnalytics, givenTrackGoogleAnalytics)
	}
}

func TestCreateCampaignRequest_GoogleAnalyticsMedium(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	expectedResult := randomCreateCampaignResult()

	expectedGoogleAnalyticsMedium := test.RandomString(12, 36)
	var givenGoogleAnalyticsMedium string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenGoogleAnalyticsMedium = req.FormValue("ga_medium")

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.CreateCampaign(req, expectedMessageID).
		GoogleAnalyticsMedium(expectedGoogleAnalyticsMedium).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedGoogleAnalyticsMedium != givenGoogleAnalyticsMedium {
		t.Fatalf(`Google Analytics medium should be "%s", "%s" given`, expectedGoogleAnalyticsMedium, givenGoogleAnalyticsMedium)
	}
}

func TestCreateCampaignRequest_GoogleAnalyticsSource(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	expectedResult := randomCreateCampaignResult()

	expectedGoogleAnalyticsSource := test.RandomString(12, 36)
	var givenGoogleAnalyticsSource string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenGoogleAnalyticsSource = req.FormValue("ga_source")

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.CreateCampaign(req, expectedMessageID).
		GoogleAnalyticsSource(expectedGoogleAnalyticsSource).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedGoogleAnalyticsSource != givenGoogleAnalyticsSource {
		t.Fatalf(`Google Analytics source should be "%s", "%s" given`, expectedGoogleAnalyticsSource, givenGoogleAnalyticsSource)
	}
}

func TestCreateCampaignRequest_GoogleAnalyticsCampaign(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	expectedResult := randomCreateCampaignResult()

	expectedGoogleAnalyticsCampaign := test.RandomString(12, 36)
	var givenGoogleAnalyticsCampaign string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenGoogleAnalyticsCampaign = req.FormValue("ga_campaign")

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.CreateCampaign(req, expectedMessageID).
		GoogleAnalyticsCampaign(expectedGoogleAnalyticsCampaign).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedGoogleAnalyticsCampaign != givenGoogleAnalyticsCampaign {
		t.Fatalf(`Google Analytics campaign should be "%s", "%s" given`, expectedGoogleAnalyticsCampaign, givenGoogleAnalyticsCampaign)
	}
}

func TestCreateCampaignRequest_GoogleAnalyticsContent(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	expectedResult := randomCreateCampaignResult()

	expectedGoogleAnalyticsContent := test.RandomString(12, 36)
	var givenGoogleAnalyticsContent string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenGoogleAnalyticsContent = req.FormValue("ga_content")

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.CreateCampaign(req, expectedMessageID).
		GoogleAnalyticsContent(expectedGoogleAnalyticsContent).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedGoogleAnalyticsContent != givenGoogleAnalyticsContent {
		t.Fatalf(`Google Analytics campaign should be "%s", "%s" given`, expectedGoogleAnalyticsContent, givenGoogleAnalyticsContent)
	}
}

func TestCreateCampaignRequest_GoogleAnalyticsTerm(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	expectedResult := randomCreateCampaignResult()

	expectedGoogleAnalyticsTerm := test.RandomString(12, 36)
	var givenGoogleAnalyticsTerm string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenGoogleAnalyticsTerm = req.FormValue("ga_term")

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.CreateCampaign(req, expectedMessageID).
		GoogleAnalyticsTerm(expectedGoogleAnalyticsTerm).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedGoogleAnalyticsTerm != givenGoogleAnalyticsTerm {
		t.Fatalf(`Google Analytics term should be "%s", "%s" given`, expectedGoogleAnalyticsTerm, givenGoogleAnalyticsTerm)
	}
}

func TestCreateCampaignRequest_Payment(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	expectedResult := randomCreateCampaignResult()

	expectedLimit := test.RandomFloat64()
	var givenLimit float64

	expectedCurrency := test.RandomString(12, 36)
	var givenCurrency string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLimit, _ = strconv.ParseFloat(req.FormValue("payment_limit"), 64)
		givenCurrency = req.FormValue("payment_currency")

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := campaigns.CreateCampaign(req, expectedMessageID).
		Payment(expectedLimit, expectedCurrency).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if math.Abs(expectedLimit-givenLimit) <= 1e-9 {
		t.Fatalf(`Limit should be %f, %f given`, expectedLimit, givenLimit)
	}

	if expectedCurrency != givenCurrency {
		t.Fatalf(`Currency should be "%s", "%s" given`, expectedCurrency, givenCurrency)
	}
}

func TestCreateCampaignRequest_Execute(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	var givenMessageID int64

	expectedResult := randomCreateCampaignResult()
	var givenResult *campaigns.CreateCampaignResult

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageID, _ = strconv.ParseInt(req.FormValue("message_id"), 10, 64)

		resp := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(resp)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := campaigns.CreateCampaign(req, expectedMessageID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageID != givenMessageID {
		t.Fatalf(`Message ID should be %d, %d given`, expectedMessageID, givenMessageID)
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal("Results should be equal")
	}
}

func randomCreateCampaignResult() *campaigns.CreateCampaignResult {
	return &campaigns.CreateCampaignResult{
		CampaignID: test.RandomInt64(9999, 999999),
		Status:     test.RandomString(12, 36),
		Count:      test.RandomInt64(9999, 999999),
	}
}
