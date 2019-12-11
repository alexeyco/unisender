package messages_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/messages"
	"github.com/alexeyco/unisender/test"
)

func TestGetTemplateRequest_SystemTemplateID(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedSystemTemplateID := test.RandomInt64(9999, 999999)
	var givenSystemTemplateID int64

	expectedResult := randomGetTemplateResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSystemTemplateID, _ = strconv.ParseInt(req.FormValue("system_template_id"), 10, 64)

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.GetTemplate(req, expectedTemplateID).
		SystemTemplateID(expectedSystemTemplateID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSystemTemplateID != givenSystemTemplateID {
		t.Fatalf(`System template ID should be %d, %d given`, expectedSystemTemplateID, givenSystemTemplateID)
	}
}

func TestGetTemplateRequest_Execute(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)
	var givenTemplateID int64

	expectedResult := randomGetTemplateResult()

	var givenResult *messages.GetTemplateResult

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTemplateID, _ = strconv.ParseInt(req.FormValue("template_id"), 10, 64)

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := messages.GetTemplate(req, expectedTemplateID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTemplateID != givenTemplateID {
		t.Fatalf(`Template ID should be %d, %d given`, expectedTemplateID, givenTemplateID)
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal("Results should be equal")
	}
}

func randomGetTemplateResult() *messages.GetTemplateResult {
	return &messages.GetTemplateResult{
		ID:                    test.RandomInt64(9999, 999999),
		SubUserLogin:          test.RandomString(12, 36),
		Title:                 test.RandomString(12, 36),
		Description:           test.RandomString(12, 36),
		Lang:                  test.RandomString(12, 36),
		Subject:               test.RandomString(12, 36),
		Attachments:           nil,
		ScreenshotURL:         test.RandomString(12, 36),
		FullsizeScreenshotURL: test.RandomString(12, 36),
		Created:               test.RandomTime(12, 365),
		MessageFormat:         test.RandomString(12, 36),
		Type:                  test.RandomString(12, 36),
		Body:                  test.RandomString(12, 36),
		BodyRaw:               test.RandomString(12, 36),
	}
}
