package contacts_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/test"
)

func TestCreateFieldRequest_TypeString(t *testing.T) {
	expectedType := "string"
	var givenType string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenType = req.FormValue("type")

		result := api.Response{
			Result: &contacts.CreateFieldResult{
				ID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.CreateField(req, test.RandomString(12, 36)).
		TypeString().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedType != givenType {
		t.Fatalf(`Type should be "%s", "%s" given`, expectedType, givenType)
	}
}

func TestCreateFieldRequest_TypeText(t *testing.T) {
	expectedType := "text"
	var givenType string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenType = req.FormValue("type")

		result := api.Response{
			Result: &contacts.CreateFieldResult{
				ID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.CreateField(req, test.RandomString(12, 36)).
		TypeText().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedType != givenType {
		t.Fatalf(`Type should be "%s", "%s" given`, expectedType, givenType)
	}
}

func TestCreateFieldRequest_TypeNumber(t *testing.T) {
	expectedType := "number"
	var givenType string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenType = req.FormValue("type")

		result := api.Response{
			Result: &contacts.CreateFieldResult{
				ID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.CreateField(req, test.RandomString(12, 36)).
		TypeNumber().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedType != givenType {
		t.Fatalf(`Type should be "%s", "%s" given`, expectedType, givenType)
	}
}

func TestCreateFieldRequest_TypeDate(t *testing.T) {
	expectedType := "date"
	var givenType string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenType = req.FormValue("type")

		result := api.Response{
			Result: &contacts.CreateFieldResult{
				ID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.CreateField(req, test.RandomString(12, 36)).
		TypeDate().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedType != givenType {
		t.Fatalf(`Type should be "%s", "%s" given`, expectedType, givenType)
	}
}

func TestCreateFieldRequest_TypeBool(t *testing.T) {
	expectedType := "bool"
	var givenType string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenType = req.FormValue("type")

		result := api.Response{
			Result: &contacts.CreateFieldResult{
				ID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.CreateField(req, test.RandomString(12, 36)).
		TypeBool().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedType != givenType {
		t.Fatalf(`Type should be "%s", "%s" given`, expectedType, givenType)
	}
}

func TestCreateFieldRequest_Execute(t *testing.T) {
	expectedName := test.RandomString(12, 36)
	var givenName string

	expectedFieldID := test.RandomInt64(9999, 999999)
	var givenFieldID int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenName = req.FormValue("name")

		result := api.Response{
			Result: &contacts.CreateFieldResult{
				ID: expectedFieldID,
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenFieldID, err := contacts.CreateField(req, expectedName).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedName != givenName {
		t.Fatalf(`Name should be "%s", "%s" given`, expectedName, givenName)
	}

	if expectedFieldID != givenFieldID {
		t.Fatalf(`Field ID should be %d, %d given`, expectedFieldID, givenFieldID)
	}
}
