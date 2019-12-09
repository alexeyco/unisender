package messages_test

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
	"github.com/alexeyco/unisender/messages"
	"github.com/alexeyco/unisender/test"
)

func TestCheckEmailRequest_Execute(t *testing.T) {
	expectedEmailIDs := test.RandomInt64Slice(12, 36)
	var givenEmailIDs []int64

	expectedResult := getRandomCheckEmailResult()
	var givenResult *messages.CheckEmailResponse

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		emailIDs := req.FormValue("email_id")
		for _, id := range strings.Split(emailIDs, ",") {
			listID, _ := strconv.ParseInt(id, 10, 64)
			givenEmailIDs = append(givenEmailIDs, listID)
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

	givenResult, err := messages.CheckEmail(req, expectedEmailIDs...).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedEmailIDs, givenEmailIDs) {
		t.Fatal("Email IDs should be equal")
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal("Request results should be equal")
	}
}

func getRandomCheckEmailResult() *messages.CheckEmailResponse {
	l := test.RandomInt(12, 32)
	statuses := make([]messages.CheckEmailResponseStatus, l)

	for i := 0; i < l; i++ {
		statuses[i] = messages.CheckEmailResponseStatus{
			ID:     test.RandomInt64(9999, 999999),
			Status: test.RandomString(12, 36),
		}
	}

	return &messages.CheckEmailResponse{
		Statuses: statuses,
	}
}
