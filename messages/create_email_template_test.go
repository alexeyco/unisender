package messages_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/messages"
	"github.com/alexeyco/unisender/test"
)

func TestCreateEmailTemplateRequest_Subject(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedSubject := test.RandomString(12, 36)
	var givenSubject string

	expectedResult := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSubject = req.FormValue("subject")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: expectedResult,
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		Subject(expectedSubject).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSubject != givenSubject {
		t.Fatalf(`Subject should be "%s", "%s" given`, expectedSubject, givenSubject)
	}
}

func TestCreateEmailTemplateRequest_Body(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedBody := test.RandomString(12, 36)
	var givenBody string

	expectedResult := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBody = req.FormValue("body")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: expectedResult,
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		Body(expectedBody).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBody != givenBody {
		t.Fatalf(`Body should be "%s", "%s" given`, expectedBody, givenBody)
	}
}

func TestCreateEmailTemplateRequest_BodyText(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedBodyText := test.RandomString(12, 36)
	var givenBodyText string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBodyText = req.FormValue("text_body")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		BodyText(expectedBodyText).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBodyText != givenBodyText {
		t.Fatalf(`Body text should be "%s", "%s" given`, expectedBodyText, givenBodyText)
	}
}

func TestCreateEmailTemplateRequest_BodyRaw(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedBodyRaw := test.RandomString(12, 36)
	var givenBodyRaw string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBodyRaw = req.FormValue("raw_body")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		BodyRaw(expectedBodyRaw).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBodyRaw != givenBodyRaw {
		t.Fatalf(`Body raw should be "%s", "%s" given`, expectedBodyRaw, givenBodyRaw)
	}
}

func TestCreateEmailTemplateRequest_LangDA(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedLang := "da"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		LangDA().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailTemplateRequest_LangDE(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedLang := "de"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		LangDE().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailTemplateRequest_LangES(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedLang := "es"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		LangES().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailTemplateRequest_LangFR(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedLang := "fr"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		LangFR().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailTemplateRequest_LangNL(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedLang := "nl"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		LangNL().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailTemplateRequest_LangPL(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedLang := "pl"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		LangPL().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailTemplateRequest_LangPT(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedLang := "pt"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		LangPT().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailTemplateRequest_LangTR(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedLang := "tr"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		LangTR().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailTemplateRequest_Description(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedDescription := test.RandomString(12, 36)
	var givenDescription string

	expectedResult := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenDescription = req.FormValue("description")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: expectedResult,
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		Description(expectedDescription).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedDescription != givenDescription {
		t.Fatalf(`Description should be "%s", "%s" given`, expectedDescription, givenDescription)
	}
}

func TestCreateEmailTemplateRequest_MessageFormatBlock(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedMessageFormat := "block"
	var givenMessageFormat string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageFormat = req.FormValue("message_format")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		MessageFormatBlock().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageFormat != givenMessageFormat {
		t.Fatalf(`Message format should be "%s", "%s" given`, expectedMessageFormat, givenMessageFormat)
	}
}

func TestCreateEmailTemplateRequest_MessageFormatRawHTML(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedMessageFormat := "raw_html"
	var givenMessageFormat string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageFormat = req.FormValue("message_format")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		MessageFormatRawHTML().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageFormat != givenMessageFormat {
		t.Fatalf(`Message format should be "%s", "%s" given`, expectedMessageFormat, givenMessageFormat)
	}
}

func TestCreateEmailTemplateRequest_MessageFormatText(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)

	expectedMessageFormat := "text"
	var givenMessageFormat string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageFormat = req.FormValue("message_format")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailTemplate(req, expectedTitle).
		MessageFormatText().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageFormat != givenMessageFormat {
		t.Fatalf(`Message format should be "%s", "%s" given`, expectedMessageFormat, givenMessageFormat)
	}
}

func TestCreateEmailTemplateRequest_Execute(t *testing.T) {
	expectedTitle := test.RandomString(9999, 999999)
	var givenTitle string

	expectedResult := test.RandomInt64(9999, 999999)
	var givenResult int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTitle = req.FormValue("title")

		result := api.Response{
			Result: &messages.CreateEmailTemplateResult{
				TemplateID: expectedResult,
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := messages.CreateEmailTemplate(req, expectedTitle).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTitle != givenTitle {
		t.Fatalf(`List ID should be "%s", "%s" given`, expectedTitle, givenTitle)
	}

	if expectedResult != givenResult {
		t.Fatalf("Template ID should be %d, %d given", expectedResult, givenResult)
	}
}
