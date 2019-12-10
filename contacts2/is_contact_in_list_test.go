package contacts2_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/contacts2"
	"github.com/alexeyco/unisender/test"
)

func TestIsContactInListRequest_ConditionOr(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)
	expectedListIDs := test.RandomInt64Slice(12, 32)

	expectedCondition := "or"
	var givenCondition string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenCondition = req.FormValue("condition")

		result := api.Response{
			Result: true,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts2.IsContactInList(req, expectedEmail, expectedListIDs...).
		ConditionOr().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedCondition != givenCondition {
		t.Fatalf(`Condition should be "%s", "%s" given`, expectedCondition, givenCondition)
	}
}

func TestIsContactInListRequest_ConditionAnd(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)
	expectedListIDs := test.RandomInt64Slice(12, 32)

	expectedCondition := "and"
	var givenCondition string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenCondition = req.FormValue("condition")

		result := api.Response{
			Result: true,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts2.IsContactInList(req, expectedEmail, expectedListIDs...).
		ConditionAnd().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedCondition != givenCondition {
		t.Fatalf(`Condition should be "%s", "%s" given`, expectedCondition, givenCondition)
	}
}

func TestIsContactInListRequest_Execute(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)
	var givenEmail string

	expectedListIDs := test.RandomInt64Slice(12, 32)
	var givenListIDs []int64

	expectedCondition := "and"
	var givenCondition string

	expectedResult := true
	var givenResult bool

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmail = req.FormValue("email")
		listIDs := req.FormValue("list_ids")
		givenCondition = req.FormValue("condition")

		for _, id := range strings.Split(listIDs, ",") {
			listID, _ := strconv.ParseInt(id, 10, 64)
			givenListIDs = append(givenListIDs, listID)
		}

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := contacts2.IsContactInList(req, expectedEmail, expectedListIDs...).
		ConditionAnd().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmail != givenEmail {
		t.Fatalf(`Email should be "%s", "%s" given`, expectedEmail, givenEmail)
	}

	if !reflect.DeepEqual(expectedListIDs, givenListIDs) {
		t.Fatal("List IDs should be equal")
	}

	if expectedCondition != givenCondition {
		t.Fatalf(`Condition should be "%s", "%s" given`, expectedCondition, givenCondition)
	}

	if expectedResult != givenResult {
		t.Fatal("Results should be equal")
	}
}
