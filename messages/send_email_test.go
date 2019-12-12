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

func TestSendEmailRequest_SenderName(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedSenderName := test.RandomString(12, 36)
	var givenSenderName string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSenderName = req.FormValue("sender_name")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		SenderName(expectedSenderName).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSenderName != givenSenderName {
		t.Fatalf(`Sender name should be "%s", "%s" given`, expectedSenderName, givenSenderName)
	}
}

func TestSendEmailRequest_SenderEmail(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedSenderEmail := test.RandomString(12, 36)
	var givenSenderEmail string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSenderEmail = req.FormValue("sender_email")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		SenderEmail(expectedSenderEmail).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSenderEmail != givenSenderEmail {
		t.Fatalf(`Sender email should be "%s", "%s" given`, expectedSenderEmail, givenSenderEmail)
	}
}

func TestSendEmailRequest_Subject(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedSubject := test.RandomString(12, 36)
	var givenSubject string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSubject = req.FormValue("subject")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		Subject(expectedSubject).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSubject != givenSubject {
		t.Fatalf(`Subject should be "%s", "%s" given`, expectedSubject, givenSubject)
	}
}

func TestSendEmailRequest_Body(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedBody := test.RandomString(12, 36)
	var givenBody string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBody = req.FormValue("body")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		Body(expectedBody).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBody != givenBody {
		t.Fatalf(`Body should be "%s", "%s" given`, expectedBody, givenBody)
	}
}

func TestSendEmailRequest_ListID(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedListID := test.RandomInt64(9999, 999999)
	var givenListID int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenListID, err = strconv.ParseInt(req.FormValue("list_id"), 10, 64)

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		ListID(expectedListID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedListID != givenListID {
		t.Fatalf(`List ID should be %d, %d given`, expectedListID, givenListID)
	}
}

func TestSendEmailRequest_Attachment(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedName := test.RandomString(12, 36)
	expectedContent := test.RandomString(12, 36)
	var givenContent string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenContent = req.FormValue(fmt.Sprintf("attachments[%s]", expectedName))

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		Attachment(expectedName, expectedContent).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedContent != givenContent {
		t.Fatalf(`Attachment content should be "%s", "%s" given`, expectedContent, givenContent)
	}
}

func TestSendEmailRequest_LangDA(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedLang := "da"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		LangDA().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Lang should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestSendEmailRequest_LangDE(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedLang := "de"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		LangDE().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Lang should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestSendEmailRequest_LangES(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedLang := "es"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		LangES().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Lang should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestSendEmailRequest_LangFR(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedLang := "fr"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		LangFR().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Lang should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestSendEmailRequest_LangNL(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedLang := "nl"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		LangNL().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Lang should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestSendEmailRequest_LangPL(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedLang := "pl"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		LangPL().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Lang should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestSendEmailRequest_LangPT(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedLang := "pt"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		LangPT().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Lang should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestSendEmailRequest_LangTR(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedLang := "tr"
	var givenLang string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLang = req.FormValue("lang")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		LangTR().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLang != givenLang {
		t.Fatalf(`Lang should be "%s", "%s" given`, expectedLang, givenLang)
	}
}

func TestSendEmailRequest_TrackRead(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedTrackRead := 1
	var givenTrackRead int

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTrackRead, _ = strconv.Atoi(req.FormValue("track_read"))

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		TrackRead().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTrackRead != givenTrackRead {
		t.Fatalf(`Track read should be %d, %d given`, expectedTrackRead, givenTrackRead)
	}
}

func TestSendEmailRequest_TrackLinks(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedTrackLinks := 1
	var givenTrackLinks int

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTrackLinks, _ = strconv.Atoi(req.FormValue("track_links"))

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		TrackLinks().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTrackLinks != givenTrackLinks {
		t.Fatalf(`Track read should be %d, %d given`, expectedTrackLinks, givenTrackLinks)
	}
}

func TestSendEmailRequest_CC(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedCC := test.RandomString(12, 36)
	var givenCC string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenCC = req.FormValue("cc")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		CC(expectedCC).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedCC != givenCC {
		t.Fatalf(`CC should be "%s", "%s" given`, expectedCC, givenCC)
	}
}

func TestSendEmailRequest_WrapTypeSkip(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedWrapType := "skip"
	var givenWrapType string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenWrapType = req.FormValue("wrap_type")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		WrapTypeSkip().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedWrapType != givenWrapType {
		t.Fatalf(`Wrap type should be "%s", "%s" given`, expectedWrapType, givenWrapType)
	}
}

func TestSendEmailRequest_WrapTypeLeft(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedWrapType := "left"
	var givenWrapType string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenWrapType = req.FormValue("wrap_type")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		WrapTypeLeft().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedWrapType != givenWrapType {
		t.Fatalf(`Wrap type should be "%s", "%s" given`, expectedWrapType, givenWrapType)
	}
}

func TestSendEmailRequest_WrapTypeRight(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedWrapType := "right"
	var givenWrapType string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenWrapType = req.FormValue("wrap_type")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		WrapTypeRight().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedWrapType != givenWrapType {
		t.Fatalf(`Wrap type should be "%s", "%s" given`, expectedWrapType, givenWrapType)
	}
}

func TestSendEmailRequest_WrapTypeCenter(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedWrapType := "center"
	var givenWrapType string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenWrapType = req.FormValue("wrap_type")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		WrapTypeCenter().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedWrapType != givenWrapType {
		t.Fatalf(`Wrap type should be "%s", "%s" given`, expectedWrapType, givenWrapType)
	}
}

func TestSendEmailRequest_ImagesAsAttachments(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedImagesAs := "attachments"
	var givenImagesAs string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenImagesAs = req.FormValue("images_as")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		ImagesAsAttachments().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedImagesAs != givenImagesAs {
		t.Fatalf(`Images as should be "%s", "%s" given`, expectedImagesAs, givenImagesAs)
	}
}

func TestSendEmailRequest_ImagesAsOnlyLinks(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedImagesAs := "only_links"
	var givenImagesAs string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenImagesAs = req.FormValue("images_as")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		ImagesAsOnlyLinks().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedImagesAs != givenImagesAs {
		t.Fatalf(`Images as should be "%s", "%s" given`, expectedImagesAs, givenImagesAs)
	}
}

func TestSendEmailRequest_ImagesAsUserDefault(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedImagesAs := "user_default"
	var givenImagesAs string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenImagesAs = req.FormValue("images_as")

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		ImagesAsUserDefault().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedImagesAs != givenImagesAs {
		t.Fatalf(`Images as should be "%s", "%s" given`, expectedImagesAs, givenImagesAs)
	}
}

func TestSendEmailRequest_RefKey(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedRefKey := test.RandomInt64(9999, 999999)
	var givenRefKey int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenRefKey, err = strconv.ParseInt(req.FormValue("ref_key"), 10, 64)

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		RefKey(expectedRefKey).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedRefKey != givenRefKey {
		t.Fatalf(`Ref key should be %d, %d given`, expectedRefKey, givenRefKey)
	}
}

func TestSendEmailRequest_MetaData(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedName := test.RandomString(12, 36)
	expectedValue := test.RandomString(12, 36)
	var givenValue string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenValue = req.FormValue(fmt.Sprintf("metadata[%s]", expectedName))

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: test.RandomInt64(9999, 999999),
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendEmail(req, expectedEmail).
		MetaData(expectedName, expectedValue).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedValue != givenValue {
		t.Fatalf(`Metadata value should be "%s", "%s" given`, expectedValue, givenValue)
	}
}

func TestSendEmailRequest_Execute(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)
	var givenEmail string

	expectedErrorChecking := 1
	var givenErrorChecking int

	expectedResult := test.RandomInt64(9999, 999999)
	var givenResult int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmail = req.FormValue("email")
		givenErrorChecking, _ = strconv.Atoi(req.FormValue("error_checking"))

		result := api.Response{
			Result: &messages.SendEmailResult{
				EmailID: expectedResult,
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := messages.SendEmail(req, expectedEmail).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmail != givenEmail {
		t.Fatalf(`Email should be "%s", "%s" given`, expectedEmail, givenEmail)
	}

	if expectedErrorChecking != givenErrorChecking {
		t.Fatalf(`Error checking should be %d, %d given`, expectedErrorChecking, givenErrorChecking)
	}

	if expectedResult != givenResult {
		t.Fatalf("Message ID should be %d, %d given", expectedResult, givenResult)
	}
}
