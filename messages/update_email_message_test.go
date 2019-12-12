package messages_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/messages"
	"github.com/alexeyco/unisender/test"
)

func TestUpdateEmailMessageRequest_SenderName(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedSenderName := test.RandomString(12, 36)
	var givenSenderName string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSenderName = req.FormValue("sender_name")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		SenderName(expectedSenderName).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSenderName != givenSenderName {
		t.Fatalf(`Sender name should be "%s", "%s" given`, expectedSenderName, givenSenderName)
	}
}

func TestUpdateEmailMessageRequest_SenderEmail(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedSenderEmail := test.RandomString(12, 36)
	var givenSenderEmail string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSenderEmail = req.FormValue("sender_email")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		SenderEmail(expectedSenderEmail).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSenderEmail != givenSenderEmail {
		t.Fatalf(`Sender email should be "%s", "%s" given`, expectedSenderEmail, givenSenderEmail)
	}
}

func TestUpdateEmailMessageRequest_Subject(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedSubject := test.RandomString(12, 36)
	var givenSubject string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSubject = req.FormValue("subject")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		Subject(expectedSubject).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSubject != givenSubject {
		t.Fatalf(`Subject should be "%s", "%s" given`, expectedSubject, givenSubject)
	}
}

func TestUpdateEmailMessageRequest_Body(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedBody := test.RandomString(12, 36)
	var givenBody string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBody = req.FormValue("body")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		Body(expectedBody).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBody != givenBody {
		t.Fatalf(`Body should be "%s", "%s" given`, expectedBody, givenBody)
	}
}

func TestUpdateEmailMessageRequest_BodyText(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedBodyText := test.RandomString(12, 36)
	var givenBodyText string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBodyText = req.FormValue("text_body")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		BodyText(expectedBodyText).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBodyText != givenBodyText {
		t.Fatalf(`Body text should be "%s", "%s" given`, expectedBodyText, givenBodyText)
	}
}

func TestUpdateEmailMessageRequest_BodyRaw(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedBodyRaw := test.RandomString(12, 36)
	var givenBodyRaw string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBodyRaw = req.FormValue("raw_body")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		BodyRaw(expectedBodyRaw).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBodyRaw != givenBodyRaw {
		t.Fatalf(`Body raw should be "%s", "%s" given`, expectedBodyRaw, givenBodyRaw)
	}
}

func TestUpdateEmailMessageRequest_ListID(t *testing.T) {

}

func TestUpdateEmailMessageRequest_MessageFormatBlock(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedMessageFormat := "block"
	var givenMessageFormat string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageFormat = req.FormValue("message_format")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		MessageFormatBlock().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageFormat != givenMessageFormat {
		t.Fatalf(`Message format should be "%s", "%s" given`, expectedMessageFormat, givenMessageFormat)
	}
}

func TestUpdateEmailMessageRequest_MessageFormatRawHTML(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedMessageFormat := "raw_html"
	var givenMessageFormat string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageFormat = req.FormValue("message_format")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		MessageFormatRawHTML().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageFormat != givenMessageFormat {
		t.Fatalf(`Message format should be "%s", "%s" given`, expectedMessageFormat, givenMessageFormat)
	}
}

func TestUpdateEmailMessageRequest_MessageFormatText(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedMessageFormat := "text"
	var givenMessageFormat string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageFormat = req.FormValue("message_format")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		MessageFormatText().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageFormat != givenMessageFormat {
		t.Fatalf(`Message format should be "%s", "%s" given`, expectedMessageFormat, givenMessageFormat)
	}
}

func TestUpdateEmailMessageRequest_LangDA(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedLang := "da"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		LangDA().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailMessageRequest_LangDE(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedLang := "de"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		LangDE().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailMessageRequest_LangES(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedLang := "es"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		LangES().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailMessageRequest_LangFR(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedLang := "fr"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		LangFR().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailMessageRequest_LangNL(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedLang := "nl"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		LangNL().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailMessageRequest_LangPL(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedLang := "pl"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		LangPL().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailMessageRequest_LangPT(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedLang := "pt"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		LangPT().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailMessageRequest_LangTR(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)

	expectedLang := "tr"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		LangTR().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailMessageRequest_Categories(t *testing.T) {

}

func TestUpdateEmailMessageRequest_Execute(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	var givenMessageID int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageID, err = strconv.ParseInt(req.FormValue("id"), 10, 64)

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailMessage(req, expectedMessageID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageID != givenMessageID {
		t.Fatalf("List ID should be %d, %d given", expectedMessageID, givenMessageID)
	}
}
