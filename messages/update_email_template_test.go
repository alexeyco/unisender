package messages_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/messages"
	"github.com/alexeyco/unisender/test"
)

func TestUpdateEmailTemplateRequest_Title(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedTitle := test.RandomString(12, 36)
	var givenTitle string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTitle = req.FormValue("title")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		Title(expectedTitle).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTitle != givenTitle {
		t.Fatalf(`Title should be "%s", "%s" given`, expectedTitle, givenTitle)
	}
}

func TestUpdateEmailTemplateRequest_Subject(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedSubject := test.RandomString(12, 36)
	var givenSubject string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSubject = req.FormValue("subject")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		Subject(expectedSubject).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSubject != givenSubject {
		t.Fatalf(`Subject should be "%s", "%s" given`, expectedSubject, givenSubject)
	}
}

func TestUpdateEmailTemplateRequest_Body(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedBody := test.RandomString(12, 36)
	var givenBody string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBody = req.FormValue("body")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		Body(expectedBody).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBody != givenBody {
		t.Fatalf(`Body should be "%s", "%s" given`, expectedBody, givenBody)
	}
}

func TestUpdateEmailTemplateRequest_BodyText(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedBodyText := test.RandomString(12, 36)
	var givenBodyText string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBodyText = req.FormValue("text_body")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		BodyText(expectedBodyText).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBodyText != givenBodyText {
		t.Fatalf(`Body text should be "%s", "%s" given`, expectedBodyText, givenBodyText)
	}
}

func TestUpdateEmailTemplateRequest_BodyRaw(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedBodyRaw := test.RandomString(12, 36)
	var givenBodyRaw string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBodyRaw = req.FormValue("raw_body")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		BodyRaw(expectedBodyRaw).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBodyRaw != givenBodyRaw {
		t.Fatalf(`Body raw should be "%s", "%s" given`, expectedBodyRaw, givenBodyRaw)
	}
}

func TestUpdateEmailTemplateRequest_LangDA(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedLang := "da"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		LangDA().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailTemplateRequest_LangDE(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedLang := "de"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		LangDE().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailTemplateRequest_LangES(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedLang := "es"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		LangES().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailTemplateRequest_LangFR(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedLang := "fr"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		LangFR().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailTemplateRequest_LangNL(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedLang := "nl"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		LangNL().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailTemplateRequest_LangPL(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedLang := "pl"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		LangPL().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailTemplateRequest_LangPT(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedLang := "pt"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		LangPT().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailTemplateRequest_LangTR(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedLang := "tr"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		LangTR().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestUpdateEmailTemplateRequest_Description(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedDescription := test.RandomString(12, 36)
	var givenDescription string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenDescription = req.FormValue("description")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		Description(expectedDescription).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedDescription != givenDescription {
		t.Fatalf(`Description should be "%s", "%s" given`, expectedDescription, givenDescription)
	}
}

func TestUpdateEmailTemplateRequest_MessageFormatBlock(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedMessageFormat := "block"
	var givenMessageFormat string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageFormat = req.FormValue("message_format")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		MessageFormatBlock().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageFormat != givenMessageFormat {
		t.Fatalf(`Message format should be "%s", "%s" given`, expectedMessageFormat, givenMessageFormat)
	}
}

func TestUpdateEmailTemplateRequest_MessageFormatRawHTML(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedMessageFormat := "raw_html"
	var givenMessageFormat string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageFormat = req.FormValue("message_format")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		MessageFormatRawHTML().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageFormat != givenMessageFormat {
		t.Fatalf(`Message format should be "%s", "%s" given`, expectedMessageFormat, givenMessageFormat)
	}
}

func TestUpdateEmailTemplateRequest_MessageFormatText(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)

	expectedMessageFormat := "text"
	var givenMessageFormat string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageFormat = req.FormValue("message_format")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		MessageFormatText().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageFormat != givenMessageFormat {
		t.Fatalf(`Message format should be "%s", "%s" given`, expectedMessageFormat, givenMessageFormat)
	}
}

func TestUpdateEmailTemplateRequest_Execute(t *testing.T) {
	expectedTemplateID := test.RandomInt64(9999, 999999)
	var givenTemplateID int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTemplateID, _ = strconv.ParseInt(req.FormValue("template_id"), 10, 64)

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.UpdateEmailTemplate(req, expectedTemplateID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTemplateID != givenTemplateID {
		t.Fatalf(`Template ID should be %d, %d given`, expectedTemplateID, givenTemplateID)
	}
}
