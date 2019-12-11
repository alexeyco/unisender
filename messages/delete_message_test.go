package messages_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/messages"
	"github.com/alexeyco/unisender/test"
)

func TestDeleteMessageRequest_Execute(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	var givenMessageID int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageID, err = strconv.ParseInt(req.FormValue("message_id"), 10, 64)

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.DeleteMessage(req, expectedMessageID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageID != givenMessageID {
		t.Fatalf("Message ID should be %d, %d given", expectedMessageID, givenMessageID)
	}
}
