package contacts_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/contacts"
)

func TestGetContactRequest_IncludeLists(t *testing.T) {
	expectedEmail := randomString(12, 48)
	expectedIncludeLists := 1
	var givenIncludeLists int

	expectedResult := randomPerson()

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenIncludeLists, _ = strconv.Atoi(req.FormValue("include_lists"))

		response, _ := json.Marshal(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.GetContact(req, expectedEmail).
		IncludeLists().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedIncludeLists != givenIncludeLists {
		t.Fatalf(`Param "include_lists" should be %d, %d given`, expectedIncludeLists, givenIncludeLists)
	}
}

func TestGetContactRequest_IncludeFields(t *testing.T) {
	expectedEmail := randomString(12, 48)
	expectedIncludeFields := 1
	var givenIncludeFields int

	expectedResult := randomPerson()

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenIncludeFields, _ = strconv.Atoi(req.FormValue("include_fields"))

		response, _ := json.Marshal(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.GetContact(req, expectedEmail).
		IncludeFields().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedIncludeFields != givenIncludeFields {
		t.Fatalf(`Param "include_fields" should be %d, %d given`, expectedIncludeFields, givenIncludeFields)
	}
}

func TestGetContactRequest_IncludeDetails(t *testing.T) {
	expectedEmail := randomString(12, 48)
	expectedIncludeDetails := 1
	var givenIncludeDetails int

	expectedResult := randomPerson()

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenIncludeDetails, _ = strconv.Atoi(req.FormValue("include_details"))

		response, _ := json.Marshal(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.GetContact(req, expectedEmail).
		IncludeDetails().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedIncludeDetails != givenIncludeDetails {
		t.Fatalf(`Param "include_details" should be %d, %d given`, expectedIncludeDetails, givenIncludeDetails)
	}
}

func TestGetContactRequest_Execute(t *testing.T) {
	expectedEmail := randomString(12, 48)
	var givenEmail string

	expectedResult := randomPerson()

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmail = req.FormValue("email")

		response, _ := json.Marshal(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := contacts.GetContact(req, expectedEmail).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmail != givenEmail {
		t.Fatalf(`Email should be "%s", "%s" given`, expectedEmail, givenEmail)
	}

	if reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal("Results should be equal")
	}
}

func randomPerson() *contacts.Person {
	var lists []contacts.PersonList
	l := randomInt(1, 12)
	for i := 0; i < l; i++ {
		lists = append(lists, contacts.PersonList{
			ID:      int64(randomInt(9999, 999999)),
			Title:   randomString(12, 36),
			AddedAt: randomTime(12, 365),
		})
	}

	fields := map[string]string{}
	l = randomInt(1, 12)
	for i := 0; i < l; i++ {
		fields[randomString(4, 16)] = randomString(12, 36)
	}

	return &contacts.Person{
		Email: contacts.PersonEmail{
			Email:        randomString(12, 36),
			AddedAt:      randomTime(12, 365),
			Status:       randomString(12, 36),
			Availability: randomString(12, 36),
			LastSend:     randomTime(12, 365),
			LastDelivery: randomTime(12, 365),
			LastRead:     randomTime(12, 365),
			LastClick:    randomTime(12, 365),
			Rating:       1.2,
			Lists:        lists,
			Fields:       fields,
		},
	}
}
