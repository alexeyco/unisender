package contacts_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/contacts"
)

func TestGetContactCountRequest_ParamsTagID(t *testing.T) {
	expectedListID := int64(randomInt(9999, 999999))
	expectedResult := randomGetContactCountResult()

	expectedParamsTagID := int64(randomInt(9999, 999999))
	var givenParamsTagID int64

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenParamsTagID, _ = strconv.ParseInt(req.FormValue("params[tagId]"), 10, 64)

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.GetContactCount(req, expectedListID).
		ParamsTagID(expectedParamsTagID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedParamsTagID != givenParamsTagID {
		t.Fatalf(`Param "tagId" ID should be %d, %d given`, expectedParamsTagID, givenParamsTagID)
	}
}

func TestGetContactCountRequest_ParamsTypeAddress(t *testing.T) {
	expectedListID := int64(randomInt(9999, 999999))
	expectedResult := randomGetContactCountResult()

	expectedParamsType := "address"
	var givenParamsType string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenParamsType = req.FormValue("params[type]")

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.GetContactCount(req, expectedListID).
		ParamsTypeAddress().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedParamsType != givenParamsType {
		t.Fatalf(`Param "type" ID should be "%s", "%s" given`, expectedParamsType, givenParamsType)
	}
}

func TestGetContactCountRequest_ParamsTypePhone(t *testing.T) {
	expectedListID := int64(randomInt(9999, 999999))
	expectedResult := randomGetContactCountResult()

	expectedParamsType := "phone"
	var givenParamsType string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenParamsType = req.FormValue("params[type]")

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.GetContactCount(req, expectedListID).
		ParamsTypePhone().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedParamsType != givenParamsType {
		t.Fatalf(`Param "type" ID should be "%s", "%s" given`, expectedParamsType, givenParamsType)
	}
}

func TestGetContactCountRequest_ParamsSearch(t *testing.T) {
	expectedListID := int64(randomInt(9999, 999999))
	expectedResult := randomGetContactCountResult()

	expectedParamsSearch := randomString(12, 48)
	var givenParamsSearch string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenParamsSearch = req.FormValue("params[search]")

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.GetContactCount(req, expectedListID).
		ParamsSearch(expectedParamsSearch).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedParamsSearch != givenParamsSearch {
		t.Fatalf(`Param "search" ID should be "%s", "%s" given`, expectedParamsSearch, givenParamsSearch)
	}
}

func TestGetContactCountRequest_Execute(t *testing.T) {
	expectedListID := int64(randomInt(9999, 999999))
	var givenListID int64

	expectedResult := randomGetContactCountResult()

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenListID, _ = strconv.ParseInt(req.FormValue("list_id"), 10, 64)

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := contacts.GetContactCount(req, expectedListID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedListID != givenListID {
		t.Fatalf(`List ID should be %d, %d given`, expectedListID, givenListID)
	}

	if expectedResult.Count != givenResult {
		t.Fatalf(`Count should be %d, %d given`, expectedResult.Count, givenResult)
	}
}

func randomGetContactCountResult() *contacts.GetContactCountResult {
	return &contacts.GetContactCountResult{
		Count: int64(randomInt(9999, 999999)),
	}
}
