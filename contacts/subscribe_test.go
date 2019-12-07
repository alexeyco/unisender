package contacts_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/test"
)

func TestSubscribeRequest_Field(t *testing.T) {
	expectedField := test.RandomString(12, 36)

	expectedValue := test.RandomString(12, 36)
	var givenValue string

	expectedListIDs := test.RandomInt64Slice(12, 36)

	expectedPersonID := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenValue = req.FormValue(fmt.Sprintf("fields[%s]", expectedField))

		result := api.Response{
			Result: &contacts.SubscribeResponse{
				PersonID: expectedPersonID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.Subscribe(req, expectedListIDs...).
		Field(expectedField, expectedValue).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedValue != givenValue {
		t.Fatalf(`Field value should be "%s", "%s" given`, expectedValue, givenValue)
	}
}

func TestSubscribeRequest_Email(t *testing.T) {
	expectedContact := test.RandomString(12, 36)
	var givenContact string

	expectedListIDs := test.RandomInt64Slice(12, 36)

	expectedPersonID := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenContact = req.FormValue("fields[email]")

		result := api.Response{
			Result: &contacts.SubscribeResponse{
				PersonID: expectedPersonID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.Subscribe(req, expectedListIDs...).
		Email(expectedContact).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedContact != givenContact {
		t.Fatalf(`Email should be "%s", "%s" given`, expectedContact, givenContact)
	}
}

func TestSubscribeRequest_Phone(t *testing.T) {
	expectedContact := test.RandomString(12, 36)
	var givenContact string

	expectedListIDs := test.RandomInt64Slice(12, 36)

	expectedPersonID := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenContact = req.FormValue("fields[phone]")

		result := api.Response{
			Result: &contacts.SubscribeResponse{
				PersonID: expectedPersonID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.Subscribe(req, expectedListIDs...).
		Phone(expectedContact).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedContact != givenContact {
		t.Fatalf(`Email should be "%s", "%s" given`, expectedContact, givenContact)
	}
}

func TestSubscribeRequest_Tags(t *testing.T) {
	expectedTags := test.RandomStringSlice(12, 32)
	var givenTags []string

	expectedListIDs := test.RandomInt64Slice(12, 36)

	expectedPersonID := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTags = strings.Split(req.FormValue("tags"), ",")

		result := api.Response{
			Result: &contacts.SubscribeResponse{
				PersonID: expectedPersonID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.Subscribe(req, expectedListIDs...).
		Tags(expectedTags...).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedTags, givenTags) {
		t.Fatal("Tags should be equal")
	}
}

func TestSubscribeRequest_DoubleOptinUnconfirmed(t *testing.T) {
	expectedContact := test.RandomString(12, 36)
	expectedListIDs := test.RandomInt64Slice(12, 36)

	expectedDoubleOptin := 0
	var givenDoubleOptin int

	expectedPersonID := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenDoubleOptin, _ = strconv.Atoi(req.FormValue("double_optin"))

		result := api.Response{
			Result: &contacts.SubscribeResponse{
				PersonID: expectedPersonID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.Subscribe(req, expectedListIDs...).
		Email(expectedContact).
		DoubleOptinUnconfirmed().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedDoubleOptin != givenDoubleOptin {
		t.Fatalf(`Param "double_optin" should be %d, %d given`, expectedDoubleOptin, givenDoubleOptin)
	}
}

func TestSubscribeRequest_DoubleOptinConfirmed(t *testing.T) {
	expectedContact := test.RandomString(12, 36)
	expectedListIDs := test.RandomInt64Slice(12, 36)

	expectedDoubleOptin := 3
	var givenDoubleOptin int

	expectedPersonID := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenDoubleOptin, _ = strconv.Atoi(req.FormValue("double_optin"))

		result := api.Response{
			Result: &contacts.SubscribeResponse{
				PersonID: expectedPersonID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.Subscribe(req, expectedListIDs...).
		Email(expectedContact).
		DoubleOptinConfirmed().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedDoubleOptin != givenDoubleOptin {
		t.Fatalf(`Param "double_optin" should be %d, %d given`, expectedDoubleOptin, givenDoubleOptin)
	}
}

func TestSubscribeRequest_DoubleOptinConfirmedIfActiveOrNew(t *testing.T) {
	expectedContact := test.RandomString(12, 36)
	expectedListIDs := test.RandomInt64Slice(12, 36)

	expectedDoubleOptin := 4
	var givenDoubleOptin int

	expectedPersonID := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenDoubleOptin, _ = strconv.Atoi(req.FormValue("double_optin"))

		result := api.Response{
			Result: &contacts.SubscribeResponse{
				PersonID: expectedPersonID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.Subscribe(req, expectedListIDs...).
		Email(expectedContact).
		DoubleOptinConfirmedIfActiveOrNew().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedDoubleOptin != givenDoubleOptin {
		t.Fatalf(`Param "double_optin" should be %d, %d given`, expectedDoubleOptin, givenDoubleOptin)
	}
}

func TestSubscribeRequest_DoNotOverwrite(t *testing.T) {
	expectedContact := test.RandomString(12, 36)
	expectedListIDs := test.RandomInt64Slice(12, 36)

	expectedOverwrite := 0
	var givenOverwrite int

	expectedPersonID := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenOverwrite, _ = strconv.Atoi(req.FormValue("overwrite"))

		result := api.Response{
			Result: &contacts.SubscribeResponse{
				PersonID: expectedPersonID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.Subscribe(req, expectedListIDs...).
		Email(expectedContact).
		DoNotOverwrite().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedOverwrite != givenOverwrite {
		t.Fatalf("Overwrite param should be %d, %d given", expectedOverwrite, givenOverwrite)
	}
}

func TestSubscribeRequest_OverwriteAll(t *testing.T) {
	expectedContact := test.RandomString(12, 36)
	expectedListIDs := test.RandomInt64Slice(12, 36)

	expectedOverwrite := 1
	var givenOverwrite int

	expectedPersonID := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenOverwrite, _ = strconv.Atoi(req.FormValue("overwrite"))

		result := api.Response{
			Result: &contacts.SubscribeResponse{
				PersonID: expectedPersonID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.Subscribe(req, expectedListIDs...).
		Email(expectedContact).
		OverwriteAll().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedOverwrite != givenOverwrite {
		t.Fatalf("Overwrite param should be %d, %d given", expectedOverwrite, givenOverwrite)
	}
}

func TestSubscribeRequest_OverwritePartially(t *testing.T) {
	expectedContact := test.RandomString(12, 36)
	expectedListIDs := test.RandomInt64Slice(12, 36)

	expectedOverwrite := 2
	var givenOverwrite int

	expectedPersonID := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenOverwrite, _ = strconv.Atoi(req.FormValue("overwrite"))

		result := api.Response{
			Result: &contacts.SubscribeResponse{
				PersonID: expectedPersonID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.Subscribe(req, expectedListIDs...).
		Email(expectedContact).
		OverwritePartially().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedOverwrite != givenOverwrite {
		t.Fatalf("Overwrite param should be %d, %d given", expectedOverwrite, givenOverwrite)
	}
}

func TestSubscribeRequest_Execute(t *testing.T) {
	expectedContact := test.RandomString(12, 36)

	expectedListIDs := test.RandomInt64Slice(12, 36)
	var givenListIDs []int64

	expectedPersonID := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		s := req.FormValue("list_ids")
		for _, idStr := range strings.Split(s, ",") {
			id, _ := strconv.ParseInt(idStr, 10, 64)
			givenListIDs = append(givenListIDs, id)
		}

		result := api.Response{
			Result: &contacts.SubscribeResponse{
				PersonID: expectedPersonID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenPersonID, err := contacts.Subscribe(req, expectedListIDs...).
		Email(expectedContact).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedListIDs, givenListIDs) {
		t.Fatal("List IDs should be equal")
	}

	if expectedPersonID != givenPersonID {
		t.Fatalf("Person ID should be %d, %d given", expectedPersonID, givenPersonID)
	}
}
