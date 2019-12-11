package messages_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/messages"
	"github.com/alexeyco/unisender/test"
)

func TestDeleteTemplateRequest_Execute(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)
	var givenTemplateID int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTemplateID, err = strconv.ParseInt(req.FormValue("template_id"), 10, 64)

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.DeleteTemplate(req, expectedTemplateID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTemplateID != givenTemplateID {
		t.Fatalf("Template ID should be %d, %d given", expectedTemplateID, givenTemplateID)
	}
}
