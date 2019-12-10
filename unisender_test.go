package unisender_test

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/alexeyco/unisender"
	"github.com/alexeyco/unisender/test"
)

func TestUniSender_ApiKey(t *testing.T) {
	apiKeyExpected := test.RandomString(12, 36)
	var apiKeyRequested string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		apiKeyRequested = req.FormValue("api_key")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(apiKeyExpected)
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if apiKeyExpected != apiKeyRequested {
		t.Fatalf(`API key should be "%s", "%s" given`, apiKeyExpected, apiKeyRequested)
	}
}

func TestUniSender_SetLanguageEnglish(t *testing.T) {
	expectedLanguage := "en"
	var givenLanguage string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		givenLanguage = req.FormValue("lang")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(test.RandomString(12, 36))
	usndr.SetLanguageEnglish()
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLanguage != givenLanguage {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLanguage, givenLanguage)
	}
}

func TestUniSender_SetLanguageItalian(t *testing.T) {
	expectedLanguage := "it"
	var givenLanguage string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		givenLanguage = req.FormValue("lang")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(test.RandomString(12, 36))
	usndr.SetLanguageItalian()
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLanguage != givenLanguage {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLanguage, givenLanguage)
	}
}

func TestUniSender_SetLanguageRussian(t *testing.T) {
	expectedLanguage := "ru"
	var givenLanguage string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		givenLanguage = req.FormValue("lang")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(test.RandomString(12, 36))
	usndr.SetLanguageRussian()
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLanguage != givenLanguage {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLanguage, givenLanguage)
	}
}

func TestUniSender_Format(t *testing.T) {
	apiKeyExpected := test.RandomString(12, 36)

	formatExpected := "json"
	var formatRequested string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		formatRequested = req.FormValue("format")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(apiKeyExpected)
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if formatExpected != formatRequested {
		t.Fatalf(`Format should be "%s", "%s" given`, formatExpected, formatRequested)
	}
}

func ExampleUniSender_GetCurrencyRates() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	currencyRates, err := usndr.GetCurrencyRates().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(currencyRates)
}

func ExampleUniSender_CancelCampaign() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	var campaignID int64 = 123
	if err := usndr.CancelCampaign(123).Execute(); err != nil {
		log.Fatalln(err)
	}

	log.Printf("Campaign (id=%d) cancelled", campaignID)
}

func ExampleUniSender_CreateCampaign() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	var messageID int64 = 123
	res, err := usndr.CreateCampaign(messageID).
		StartTime(time.Now().Add(3 * time.Hour)).
		TrackRead().
		TrackLinks().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetCampaignCommonStats() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	var campaignID int64 = 123
	res, err := usndr.GetCampaignCommonStats(campaignID).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetCampaigns() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetCampaigns().
		From(time.Now().Add(-time.Hour)).
		To(time.Now()).
		Limit(100).
		Offset(0).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}
