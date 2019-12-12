package unisender_test

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/alexeyco/unisender"
	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/test"
)

func TestUniSender_ApiKey(t *testing.T) {
	apiKeyExpected := test.RandomString(12, 36)
	var apiKeyRequested string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		apiKeyRequested = req.FormValue("api_key")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(apiKeyExpected)
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if apiKeyExpected != apiKeyRequested {
		t.Fatalf(`API key should be "%s", "%s" given`, apiKeyExpected, apiKeyRequested)
	}
}

func TestUniSender_SetLanguageEnglish(t *testing.T) {
	expectedLanguage := "en"
	var givenLanguage string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		givenLanguage = req.FormValue("lang")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(test.RandomString(12, 36))
	usndr.SetLanguageEnglish()
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLanguage != givenLanguage {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLanguage, givenLanguage)
	}
}

func TestUniSender_SetLanguageItalian(t *testing.T) {
	expectedLanguage := "it"
	var givenLanguage string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		givenLanguage = req.FormValue("lang")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(test.RandomString(12, 36))
	usndr.SetLanguageItalian()
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLanguage != givenLanguage {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLanguage, givenLanguage)
	}
}

func TestUniSender_SetLanguageRussian(t *testing.T) {
	expectedLanguage := "ru"
	var givenLanguage string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		givenLanguage = req.FormValue("lang")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(test.RandomString(12, 36))
	usndr.SetLanguageRussian()
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLanguage != givenLanguage {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLanguage, givenLanguage)
	}
}

func TestUniSender_Format(t *testing.T) {
	apiKeyExpected := test.RandomString(12, 36)

	formatExpected := "json"
	var formatRequested string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		formatRequested = req.FormValue("format")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(apiKeyExpected)
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if formatExpected != formatRequested {
		t.Fatalf(`Format should be "%s", "%s" given`, formatExpected, formatRequested)
	}
}

type caller func(usndr *unisender.UniSender)

func TestUnisender(t *testing.T) {
	data := map[string]caller{
		// Campaigns
		"cancelCampaign": func(usndr *unisender.UniSender) {
			_ = usndr.CancelCampaign(0).Execute()
		},
		"createCampaign": func(usndr *unisender.UniSender) {
			_, _ = usndr.CreateCampaign(0).Execute()
		},
		"getCampaignCommonStats": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetCampaignCommonStats(0).Execute()
		},
		"getCampaignStatus": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetCampaignStatus(0).Execute()
		},
		"getCampaigns": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetCampaigns().Execute()
		},
		"getVisitedLinks": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetVisitedLinks(0).Execute()
		},
		"getWebVersion": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetWebVersion(0).Execute()
		},
		// Common
		"getCurrencyRates": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetCurrencyRates().Execute()
		},
		// Contacts
		"createField": func(usndr *unisender.UniSender) {
			_, _ = usndr.CreateField("").Execute()
		},
		"deleteField": func(usndr *unisender.UniSender) {
			_ = usndr.DeleteField(0).Execute()
		},
		"deleteTag": func(usndr *unisender.UniSender) {
			_ = usndr.DeleteTag(0).Execute()
		},
		"exclude": func(usndr *unisender.UniSender) {
			_ = usndr.Exclude("").Execute()
		},
		"exportContacts": func(usndr *unisender.UniSender) {
			_, _ = usndr.ExportContacts().Execute()
		},
		"getContact": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetContact("").Execute()
		},
		"getContactCount": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetContactCount(0).Execute()
		},
		"getContactFieldValues": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetContactFieldValues("").Execute()
		},
		"getFields": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetFields().Execute()
		},
		"getTags": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetTags().Execute()
		},
		"getTotalContactsCount": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetTotalContactsCount("").Execute()
		},
		"importContacts": func(usndr *unisender.UniSender) {
			_, _ = usndr.ImportContacts(contacts.NewImportContactsCollection()).Execute()
		},
		"isContactInList": func(usndr *unisender.UniSender) {
			_, _ = usndr.IsContactInList("").Execute()
		},
		"subscribe": func(usndr *unisender.UniSender) {
			_, _ = usndr.Subscribe().Execute()
		},
		"unsubscribe": func(usndr *unisender.UniSender) {
			_ = usndr.Unsubscribe("").Execute()
		},
		"updateField": func(usndr *unisender.UniSender) {
			_, _ = usndr.UpdateField(0, "").Execute()
		},
		// Lists
		"createList": func(usndr *unisender.UniSender) {
			_, _ = usndr.CreateList("").Execute()
		},
		"deleteList": func(usndr *unisender.UniSender) {
			_ = usndr.DeleteList(0).Execute()
		},
		"getLists": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetLists().Execute()
		},
		"updateList": func(usndr *unisender.UniSender) {
			_ = usndr.UpdateList(0, "").Execute()
		},
		"updateOptInEmail": func(usndr *unisender.UniSender) {
			_ = usndr.UpdateOptInEmail(0).Execute()
		},
		// Messages
		"checkEmail": func(usndr *unisender.UniSender) {
			_, _ = usndr.CheckEmail().Execute()
		},
		"checkSms": func(usndr *unisender.UniSender) {
			_, _ = usndr.CheckSMS(0).Execute()
		},
		"createEmailMessage": func(usndr *unisender.UniSender) {
			_, _ = usndr.CreateEmailMessage(0).Execute()
		},
		"createEmailTemplate": func(usndr *unisender.UniSender) {
			_, _ = usndr.CreateEmailTemplate("").Execute()
		},
		"createSmsMessage": func(usndr *unisender.UniSender) {
			_, _ = usndr.CreateSMSMessage("").Execute()
		},
		"deleteMessage": func(usndr *unisender.UniSender) {
			_ = usndr.DeleteMessage(0).Execute()
		},
		"deleteTemplate": func(usndr *unisender.UniSender) {
			_ = usndr.DeleteTemplate(0).Execute()
		},
		"getActualMessageVersion": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetActualMessageVersion(0).Execute()
		},
		"getMessage": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetMessage(0).Execute()
		},
		"getMessages": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetMessages().Execute()
		},
		"getTemplate": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetTemplate(0).Execute()
		},
		"getTemplates": func(usndr *unisender.UniSender) {
			_, _ = usndr.GetTemplates().Execute()
		},
		"listMessages": func(usndr *unisender.UniSender) {
			_, _ = usndr.ListMessages().Execute()
		},
		"listTemplates": func(usndr *unisender.UniSender) {
			_, _ = usndr.ListTemplates().Execute()
		},
		"sendEmail": func(usndr *unisender.UniSender) {
			_, _ = usndr.SendEmail("").Execute()
		},
		"sendSms": func(usndr *unisender.UniSender) {
			_, _ = usndr.SendSMS("").Execute()
		},
		"sendTestEmail": func(usndr *unisender.UniSender) {
			_ = usndr.SendTestEmail(0).Execute()
		},
		"updateEmailMessage": func(usndr *unisender.UniSender) {
			_ = usndr.UpdateEmailMessage(0).Execute()
		},
		"updateEmailTemplate": func(usndr *unisender.UniSender) {
			_ = usndr.UpdateEmailTemplate(0).Execute()
		},
		// Partners
	}

	var givenUrl string
	usndr := unisender.New("nothing").
		SetLanguageEnglish().
		SetClient(test.NewClient(func(req *http.Request) (res *http.Response, err error) {
			givenUrl = req.URL.String()

			res = &http.Response{
				StatusCode: http.StatusOK,
			}

			return
		}))

	for method, fn := range data {
		expectedUrl := fmt.Sprintf("https://api.unisender.com/en/api/%s", method)
		fn(usndr)

		if expectedUrl != givenUrl {
			t.Errorf(`URL should be "%s", "%s" given`, expectedUrl, givenUrl)
		}
	}
}

func ExampleUniSender_GetCurrencyRates() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	currencyRates, err := usndr.GetCurrencyRates().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(currencyRates)
}

func ExampleUniSender_CancelCampaign() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	if err := usndr.CancelCampaign(123).Execute(); err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_CreateCampaign() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.CreateCampaign(123).
		StartTime(time.Now().Add(3 * time.Hour)).
		TrackRead().
		TrackLinks().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetCampaignCommonStats() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetCampaignCommonStats(123).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetCampaigns() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetCampaigns().
		From(time.Now().Add(-time.Hour)).
		To(time.Now()).
		Limit(100).
		Offset(0).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetVisitedLinks() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetVisitedLinks(123).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetWebVersion() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetWebVersion(123).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_CreateField() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	fieldID, err := usndr.CreateField("SomeField").
		TypeString().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(fieldID)
}

func ExampleUniSender_DeleteField() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	err := usndr.DeleteField(123).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_DeleteTag() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	err := usndr.DeleteTag(123).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_Exclude() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	email := "foo@bar.example"
	err := usndr.Exclude(email).
		ContactTypeEmail().
		ListIDs(1, 2, 3).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_ExportContacts() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.ExportContacts().
		EmailStatusBlocked().
		NotifyUrl("https://foo.bar/example").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetContact() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetContact("foo@bar.example").
		IncludeLists().
		IncludeFields().
		IncludeDetails().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetContactCount() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetContactCount(1).
		ParamsTypeAddress().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetContactFieldValues() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetContactFieldValues("foo@bar.example", 1, 2, 3).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetFields() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetFields().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetTags() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetTags().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetTotalContactsCount() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetTotalContactsCount("my-login").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_ImportContacts() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	collection := contacts.NewImportContactsCollection()
	collection.Email("foo1@bar.example").
		AddListID(1, time.Now()).
		SetAvailabilityAvailable().
		SetStatusActive()

	collection.Email("foo2@bar.example").
		Delete()

	res, err := usndr.ImportContacts(collection).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_IsContactInList() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.IsContactInList("foo@bar.example", 1, 2, 3).
		ConditionOr().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_Subscribe() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.Subscribe().
		Email("foo@bar.example").
		Tags("foo", "bar", "example").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_Unsubscribe() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	err := usndr.Unsubscribe("foo@bar.example").
		ContactTypeEmail().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_UpdateField() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.UpdateField(123, "Name").
		PublicName("PublicName").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_CreateList() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.CreateList("My new mailing list").
		BeforeSubscribeUrl("https://before-subscribe.url").
		AfterSubscribeUrl("https://after-subscribe.url").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_DeleteList() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	err := usndr.DeleteList(123).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_GetLists() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetLists().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_UpdateList() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	err := usndr.UpdateList(123, "Mailing list new name").
		BeforeSubscribeUrl("https://before-subscribe.url").
		AfterSubscribeUrl("https://after-subscribe.url").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_UpdateOptInEmail() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	err := usndr.UpdateOptInEmail(123).
		SenderName("John Doe").
		SenderEmail("foo@bar.example").
		Subject("Welcome aboard!").
		Body("<b>Hi!</b>").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_CheckEmail() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.CheckEmail(123).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_CheckSMS() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.CheckSMS(123).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_CreateEmailMessage() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.CreateEmailMessage(123).
		SenderName("John Doe").
		SenderEmail("foo@bar.example").
		Subject("Welcome aboard!").
		Body("<b>Hi!</b>").
		GenerateText().
		WrapTypeSkip().
		LangDE().
		Tag("foo").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_CreateEmailTemplate() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.CreateEmailTemplate("My template title").
		Subject("Welcome aboard!").
		Body("<b>Hi!</b>").
		LangDE().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_CreateSMSMessage() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.CreateSMSMessage("+1234567").
		Body("<b>Hi!</b>").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_DeleteMessage() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	err := usndr.DeleteMessage(123).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_DeleteTemplate() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	err := usndr.DeleteTemplate(123).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_GetActualMessageVersion() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetActualMessageVersion(123).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetMessages() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetMessages().
		From(time.Now().Add(-24 * time.Hour)).
		To(time.Now()).
		Limit(30).
		Offset(0).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetMessage() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetMessage(123).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetTemplates() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetTemplates().
		TypeSystem().
		From(time.Now().Add(-24 * time.Hour)).
		To(time.Now()).
		Limit(30).
		Offset(0).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetTemplate() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetTemplate(123).
		SystemTemplateID(456).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_ListMessages() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.ListMessages().
		From(time.Now().Add(-24 * time.Hour)).
		To(time.Now()).
		Limit(30).
		Offset(0).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_ListTemplates() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.ListTemplates().
		TypeSystem().
		From(time.Now().Add(-24 * time.Hour)).
		To(time.Now()).
		Limit(30).
		Offset(0).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_SendEmail() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.SendEmail("foo@bar.example").
		SenderName("John Doe").
		SenderEmail("john@doe.example").
		Subject("Hi there").
		Body("<p>Hi there!</p>").
		LangFR().
		WrapTypeSkip().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_SendSMS() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.SendSMS("+1234567890").
		Sender("+987654321").
		Text("Hi there!").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_SendTestEmail() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	err := usndr.SendTestEmail(123).
		To("foo@bar.example").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_UpdateEmailMessage() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	err := usndr.UpdateEmailMessage(123).
		SenderName("John Doe").
		SenderEmail("foo@bar.example").
		Subject("Welcome aboard!").
		Body("<b>Hi!</b>").
		ListID(456).
		LangDE().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_UpdateEmailTemplate() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	err := usndr.UpdateEmailTemplate(123).
		Title("New template title").
		Subject("Welcome aboard!").
		Body("<b>Hi!</b>").
		LangDE().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_GetCheckedEmail() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetCheckedEmail("my-login").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetSenderDomainList() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetSenderDomainList("my-login").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_SetSenderDomain() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.SetSenderDomain("my-login", "john-doe.com").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_ValidateSender() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.ValidateSender("foo@bar.example").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}
