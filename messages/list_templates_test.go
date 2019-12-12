package messages_test

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
	"github.com/alexeyco/unisender/messages"
	"github.com/alexeyco/unisender/test"
)

func TestListTemplatesRequest_TypeUser(t *testing.T) {
	expectedType := "user"
	var givenType string

	expectedResult := randomListTemplatesResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenType = req.FormValue("type")

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.ListTemplates(req).
		TypeUser().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedType != givenType {
		t.Fatalf(`Type should be "%s", "%s" given`, expectedType, givenType)
	}
}

func TestListTemplatesRequest_TypeSystem(t *testing.T) {
	expectedType := "system"
	var givenType string

	expectedResult := randomListTemplatesResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenType = req.FormValue("type")

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.ListTemplates(req).
		TypeSystem().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedType != givenType {
		t.Fatalf(`Type should be "%s", "%s" given`, expectedType, givenType)
	}
}

func TestListTemplatesRequest_From(t *testing.T) {
	expectedFrom := test.RandomTime(12, 365)
	var givenFrom time.Time

	expectedResult := randomListTemplatesResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenFrom, _ = time.Parse("2006-01-02 15:04:05", req.FormValue("date_from"))

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.ListTemplates(req).
		From(expectedFrom).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedFrom != givenFrom {
		t.Fatalf(`From should be "%s", "%s" given`, expectedFrom, givenFrom)
	}
}

func TestListTemplatesRequest_To(t *testing.T) {
	expectedTo := test.RandomTime(12, 365)
	var givenTo time.Time

	expectedResult := randomListTemplatesResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTo, _ = time.Parse("2006-01-02 15:04:05", req.FormValue("date_to"))

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.ListTemplates(req).
		To(expectedTo).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTo != givenTo {
		t.Fatalf(`To should be "%s", "%s" given`, expectedTo, givenTo)
	}
}

func TestListTemplatesRequest_Limit(t *testing.T) {
	expectedLimit := test.RandomInt(9999, 999999)
	var givenLimit int

	expectedResult := randomListTemplatesResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLimit, _ = strconv.Atoi(req.FormValue("limit"))

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.ListTemplates(req).
		Limit(expectedLimit).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLimit != givenLimit {
		t.Fatalf(`Limit should be %d, %d given`, expectedLimit, givenLimit)
	}
}

func TestListTemplatesRequest_Offset(t *testing.T) {
	expectedOffset := test.RandomInt(9999, 999999)
	var givenOffset int

	expectedResult := randomListTemplatesResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenOffset, _ = strconv.Atoi(req.FormValue("offset"))

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.ListTemplates(req).
		Offset(expectedOffset).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedOffset != givenOffset {
		t.Fatalf(`Offset should be %d, %d given`, expectedOffset, givenOffset)
	}
}

func TestListTemplatesRequest_Execute(t *testing.T) {
	expectedResult := randomListTemplatesResult()
	var givenResult []messages.ListTemplatesResult

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := messages.ListTemplates(req).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal("Results should be equal")
	}
}

func randomListTemplatesResult() (res []messages.ListTemplatesResult) {
	num := test.RandomInt(12, 36)
	res = make([]messages.ListTemplatesResult, num)

	for i := 0; i < num; i++ {
		res[i] = messages.ListTemplatesResult{
			ID:                    test.RandomInt64(9999, 999999),
			SubUserLogin:          test.RandomString(12, 36),
			Title:                 test.RandomString(12, 36),
			Description:           test.RandomString(12, 36),
			Lang:                  test.RandomString(12, 36),
			Subject:               test.RandomString(12, 36),
			Attachments:           nil,
			ScreenshotURL:         test.RandomString(12, 36),
			Created:               test.RandomTime(12, 365),
			MessageFormat:         test.RandomString(12, 36),
			Type:                  test.RandomString(12, 36),
			FullsizeScreenshotURL: test.RandomString(12, 36),
		}
	}

	return
}
