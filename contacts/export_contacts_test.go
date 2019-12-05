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

	"github.com/alexeyco/unisender/contacts"
)

func TestExportContactsRequest_NotifyUrl(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedNotifyUrl := "https://notify-url.com"
	var givenNotifyUrl string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenNotifyUrl = req.FormValue("notify_url")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		NotifyUrl(expectedNotifyUrl).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedNotifyUrl != givenNotifyUrl {
		t.Fatalf(`Notify URL should be "%s", "%s" given`, expectedNotifyUrl, givenNotifyUrl)
	}
}

func TestExportContactsRequest_ListID(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedListID := int64(randomInt(9999, 999999))
	var givenListID int64

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenListID, _ = strconv.ParseInt(req.FormValue("list_id"), 10, 64)

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		ListID(expectedListID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedListID != givenListID {
		t.Fatalf(`List ID should be %d, %d given`, expectedListID, givenListID)
	}
}

func TestExportContactsRequest_FieldNames(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedFieldNames := randomStringSlice(16, 32)
	var givenFieldNames []string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenFieldNames = strings.Split(req.FormValue("field_names"), ",")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		FieldNames(expectedFieldNames...).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if reflect.DeepEqual(expectedFieldNames, givenFieldNames) {
		t.Fatal("Field names should be equal")
	}
}

func TestExportContactsRequest_Email(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedEmail := "foo@bar.example"
	var givenEmail string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmail = req.FormValue("email")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		Email(expectedEmail).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmail != givenEmail {
		t.Fatalf(`Email should be "%s", "%s" given`, expectedEmail, givenEmail)
	}
}

func TestExportContactsRequest_Phone(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedPhone := "+1234567890"
	var givenPhone string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenPhone = req.FormValue("phone")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		Phone(expectedPhone).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedPhone != givenPhone {
		t.Fatalf(`Phone should be "%s", "%s" given`, expectedPhone, givenPhone)
	}
}

func TestExportContactsRequest_Tag(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedTag := fmt.Sprintf("%d", randomInt(9999, 999999))
	var givenTag string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTag = req.FormValue("tag")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		Tag(expectedTag).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTag != givenTag {
		t.Fatalf(`Tag should be "%s", "%s" given`, expectedTag, givenTag)
	}
}

func TestExportContactsRequest_EmailStatusNew(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedEmailStatus := "new"
	var givenEmailStatus string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmailStatus = req.FormValue("email_status")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		EmailStatusNew().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmailStatus != givenEmailStatus {
		t.Fatalf(`Email status should be "%s", "%s" given`, expectedEmailStatus, givenEmailStatus)
	}
}

func TestExportContactsRequest_EmailStatusInvited(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedEmailStatus := "invited"
	var givenEmailStatus string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmailStatus = req.FormValue("email_status")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		EmailStatusInvited().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmailStatus != givenEmailStatus {
		t.Fatalf(`Email status should be "%s", "%s" given`, expectedEmailStatus, givenEmailStatus)
	}
}

func TestExportContactsRequest_EmailStatusActive(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedEmailStatus := "active"
	var givenEmailStatus string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmailStatus = req.FormValue("email_status")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		EmailStatusActive().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmailStatus != givenEmailStatus {
		t.Fatalf(`Email status should be "%s", "%s" given`, expectedEmailStatus, givenEmailStatus)
	}
}

func TestExportContactsRequest_EmailStatusInactive(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedEmailStatus := "inactive"
	var givenEmailStatus string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmailStatus = req.FormValue("email_status")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		EmailStatusInactive().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmailStatus != givenEmailStatus {
		t.Fatalf(`Email status should be "%s", "%s" given`, expectedEmailStatus, givenEmailStatus)
	}
}

func TestExportContactsRequest_EmailStatusUnsubscribed(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedEmailStatus := "unsubscribed"
	var givenEmailStatus string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmailStatus = req.FormValue("email_status")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		EmailStatusUnsubscribed().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmailStatus != givenEmailStatus {
		t.Fatalf(`Email status should be "%s", "%s" given`, expectedEmailStatus, givenEmailStatus)
	}
}

func TestExportContactsRequest_EmailStatusBlocked(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedEmailStatus := "blocked"
	var givenEmailStatus string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmailStatus = req.FormValue("email_status")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		EmailStatusBlocked().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmailStatus != givenEmailStatus {
		t.Fatalf(`Email status should be "%s", "%s" given`, expectedEmailStatus, givenEmailStatus)
	}
}

func TestExportContactsRequest_EmailStatusActivationRequested(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedEmailStatus := "activation_requested"
	var givenEmailStatus string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmailStatus = req.FormValue("email_status")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		EmailStatusActivationRequested().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmailStatus != givenEmailStatus {
		t.Fatalf(`Email status should be "%s", "%s" given`, expectedEmailStatus, givenEmailStatus)
	}
}

func TestExportContactsRequest_PhoneStatusNew(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedPhoneStatus := "new"
	var givenPhoneStatus string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenPhoneStatus = req.FormValue("phone_status")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		PhoneStatusNew().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedPhoneStatus != givenPhoneStatus {
		t.Fatalf(`Phone status should be "%s", "%s" given`, expectedPhoneStatus, givenPhoneStatus)
	}
}

func TestExportContactsRequest_PhoneStatusActive(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedPhoneStatus := "active"
	var givenPhoneStatus string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenPhoneStatus = req.FormValue("phone_status")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		PhoneStatusActive().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedPhoneStatus != givenPhoneStatus {
		t.Fatalf(`Phone status should be "%s", "%s" given`, expectedPhoneStatus, givenPhoneStatus)
	}
}

func TestExportContactsRequest_PhoneStatusInactive(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedPhoneStatus := "inactive"
	var givenPhoneStatus string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenPhoneStatus = req.FormValue("phone_status")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		PhoneStatusInactive().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedPhoneStatus != givenPhoneStatus {
		t.Fatalf(`Phone status should be "%s", "%s" given`, expectedPhoneStatus, givenPhoneStatus)
	}
}

func TestExportContactsRequest_PhoneStatusUnsubscribed(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedPhoneStatus := "unsubscribed"
	var givenPhoneStatus string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenPhoneStatus = req.FormValue("phone_status")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		PhoneStatusUnsubscribed().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedPhoneStatus != givenPhoneStatus {
		t.Fatalf(`Phone status should be "%s", "%s" given`, expectedPhoneStatus, givenPhoneStatus)
	}
}

func TestExportContactsRequest_PhoneStatusBlocked(t *testing.T) {
	expectedResult := randomExportContactsResult()

	expectedPhoneStatus := "blocked"
	var givenPhoneStatus string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenPhoneStatus = req.FormValue("phone_status")

		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.ExportContacts(req).
		PhoneStatusBlocked().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedPhoneStatus != givenPhoneStatus {
		t.Fatalf(`Phone status should be "%s", "%s" given`, expectedPhoneStatus, givenPhoneStatus)
	}
}

func TestExportContactsRequest_Execute(t *testing.T) {
	expectedResult := randomExportContactsResult()

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		response := exportContactResultToJson(expectedResult)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	givenResult, err := contacts.ExportContacts(req).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal("Results should be equal")
	}
}

func randomExportContactsResult() *contacts.ExportContactResult {
	return &contacts.ExportContactResult{
		TaskUUID: fmt.Sprintf("%d", randomInt(9999, 9999999)),
		Status:   fmt.Sprintf("%d", randomInt(9999, 9999999)),
	}
}

func exportContactResultToJson(res *contacts.ExportContactResult) string {
	b, _ := json.Marshal(res)
	return string(b)
}

func randomStringSlice(min, max int) []string {
	l := randomInt(min, max)
	slice := make([]string, l)
	for i := 0; i < l; i++ {
		slice[i] = fmt.Sprintf("%d", randomInt(9999, 999999))
	}

	return slice
}
