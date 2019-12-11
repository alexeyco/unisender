package messages_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/messages"
	"github.com/alexeyco/unisender/test"
)

func TestCreateEmailMessageRequest_SenderName(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedSenderName := test.RandomString(12, 36)
	var givenSenderName string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSenderName = req.FormValue("sender_name")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		SenderName(expectedSenderName).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSenderName != givenSenderName {
		t.Fatalf(`Sender name should be "%s", "%s" given`, expectedSenderName, givenSenderName)
	}
}

func TestCreateEmailMessageRequest_SenderEmail(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedSenderEmail := test.RandomString(12, 36)
	var givenSenderEmail string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSenderEmail = req.FormValue("sender_email")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		SenderEmail(expectedSenderEmail).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSenderEmail != givenSenderEmail {
		t.Fatalf(`Sender email should be "%s", "%s" given`, expectedSenderEmail, givenSenderEmail)
	}
}

func TestCreateEmailMessageRequest_Subject(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedSubject := test.RandomString(12, 36)
	var givenSubject string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSubject = req.FormValue("subject")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		Subject(expectedSubject).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSubject != givenSubject {
		t.Fatalf(`Subject should be "%s", "%s" given`, expectedSubject, givenSubject)
	}
}

func TestCreateEmailMessageRequest_Body(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedBody := test.RandomString(12, 36)
	var givenBody string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBody = req.FormValue("body")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		Body(expectedBody).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBody != givenBody {
		t.Fatalf(`Body should be "%s", "%s" given`, expectedBody, givenBody)
	}
}

func TestCreateEmailMessageRequest_BodyText(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedBodyText := test.RandomString(12, 36)
	var givenBodyText string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBodyText = req.FormValue("text_body")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		BodyText(expectedBodyText).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBodyText != givenBodyText {
		t.Fatalf(`Body text should be "%s", "%s" given`, expectedBodyText, givenBodyText)
	}
}

func TestCreateEmailMessageRequest_BodyRaw(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedBodyRaw := test.RandomString(12, 36)
	var givenBodyRaw string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBodyRaw = req.FormValue("raw_body")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		BodyRaw(expectedBodyRaw).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBodyRaw != givenBodyRaw {
		t.Fatalf(`Body raw should be "%s", "%s" given`, expectedBodyRaw, givenBodyRaw)
	}
}

func TestCreateEmailMessageRequest_GenerateText(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedGenerateText := 1
	var givenGenerateText int

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenGenerateText, _ = strconv.Atoi(req.FormValue("generate_text"))

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		GenerateText().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedGenerateText != givenGenerateText {
		t.Fatalf(`Param "generate_text" should be %d, %d given`, expectedGenerateText, givenGenerateText)
	}
}

func TestCreateEmailMessageRequest_MessageFormatBlock(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedMessageFormat := "block"
	var givenMessageFormat string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageFormat = req.FormValue("message_format")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		MessageFormatBlock().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageFormat != givenMessageFormat {
		t.Fatalf(`Message format should be "%s", "%s" given`, expectedMessageFormat, givenMessageFormat)
	}
}

func TestCreateEmailMessageRequest_MessageFormatRawHTML(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedMessageFormat := "raw_html"
	var givenMessageFormat string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageFormat = req.FormValue("message_format")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		MessageFormatRawHTML().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageFormat != givenMessageFormat {
		t.Fatalf(`Message format should be "%s", "%s" given`, expectedMessageFormat, givenMessageFormat)
	}
}

func TestCreateEmailMessageRequest_MessageFormatText(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedMessageFormat := "text"
	var givenMessageFormat string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageFormat = req.FormValue("message_format")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		MessageFormatText().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageFormat != givenMessageFormat {
		t.Fatalf(`Message format should be "%s", "%s" given`, expectedMessageFormat, givenMessageFormat)
	}
}

func TestCreateEmailMessageRequest_Tag(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedTag := test.RandomString(12, 36)
	var givenTag string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTag = req.FormValue("tag")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		Tag(expectedTag).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTag != givenTag {
		t.Fatalf(`Tag should be "%s", "%s" given`, expectedTag, givenTag)
	}
}

func TestCreateEmailMessageRequest_Attachment(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedName := test.RandomString(12, 36)
	expectedContent := test.RandomString(12, 36)
	var givenContent string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenContent = req.FormValue(fmt.Sprintf("attachments[%s]", expectedName))

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		Attachment(expectedName, expectedContent).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedContent != givenContent {
		t.Fatalf(`Content should be "%s", "%s" given`, expectedContent, givenContent)
	}
}

func TestCreateEmailMessageRequest_LangDA(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedLang := "da"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		LangDA().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailMessageRequest_LangDE(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedLang := "de"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		LangDE().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailMessageRequest_LangES(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedLang := "es"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		LangES().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailMessageRequest_LangFR(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedLang := "fr"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		LangFR().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailMessageRequest_LangNL(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedLang := "nl"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		LangNL().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailMessageRequest_LangPL(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedLang := "pl"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		LangPL().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailMessageRequest_LangPT(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedLang := "pt"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		LangPT().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailMessageRequest_LangTR(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedLang := "tr"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		LangTR().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestCreateEmailMessageRequest_TemplateID(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedTemplateID := test.RandomInt64(9999, 999999)
	var givenTemplateID int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTemplateID, err = strconv.ParseInt(req.FormValue("template_id"), 10, 64)

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		TemplateID(expectedTemplateID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTemplateID != givenTemplateID {
		t.Fatalf("Template ID should be %d, %d given", expectedTemplateID, givenTemplateID)
	}
}

func TestCreateEmailMessageRequest_SystemTemplateID(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedSystemTemplateID := test.RandomInt64(9999, 999999)
	var givenSystemTemplateID int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSystemTemplateID, err = strconv.ParseInt(req.FormValue("system_template_id"), 10, 64)

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		SystemTemplateID(expectedSystemTemplateID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSystemTemplateID != givenSystemTemplateID {
		t.Fatalf("System template ID should be %d, %d given", expectedSystemTemplateID, givenSystemTemplateID)
	}
}

func TestCreateEmailMessageRequest_WrapTypeSkip(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedWrapType := "skip"
	var givenWrapType string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenWrapType = req.FormValue("wrap_type")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		WrapTypeSkip().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedWrapType != givenWrapType {
		t.Fatalf(`Wrap type should be "%s", "%s" given`, expectedWrapType, givenWrapType)
	}
}

func TestCreateEmailMessageRequest_WrapTypeRight(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedWrapType := "right"
	var givenWrapType string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenWrapType = req.FormValue("wrap_type")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		WrapTypeRight().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedWrapType != givenWrapType {
		t.Fatalf(`Wrap type should be "%s", "%s" given`, expectedWrapType, givenWrapType)
	}
}

func TestCreateEmailMessageRequest_WrapTypeLeft(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedWrapType := "left"
	var givenWrapType string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenWrapType = req.FormValue("wrap_type")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		WrapTypeLeft().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedWrapType != givenWrapType {
		t.Fatalf(`Wrap type should be "%s", "%s" given`, expectedWrapType, givenWrapType)
	}
}

func TestCreateEmailMessageRequest_WrapTypeCenter(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedWrapType := "center"
	var givenWrapType string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenWrapType = req.FormValue("wrap_type")

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateEmailMessage(req, expectedListID).
		WrapTypeCenter().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedWrapType != givenWrapType {
		t.Fatalf(`Wrap type should be "%s", "%s" given`, expectedWrapType, givenWrapType)
	}
}

func TestCreateEmailMessageRequest_Execute(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)
	var givenListID int64

	expectedResult := test.RandomInt64(9999, 999999)
	var givenResult int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenListID, err = strconv.ParseInt(req.FormValue("list_id"), 10, 64)

		result := api.Response{
			Result: &messages.CreateEmailMessageResult{
				MessageID: expectedResult,
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := messages.CreateEmailMessage(req, expectedListID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedListID != givenListID {
		t.Fatalf("List ID should be %d, %d given", expectedListID, givenListID)
	}

	if expectedResult != givenResult {
		t.Fatalf("Message ID should be %d, %d given", expectedResult, givenResult)
	}
}
