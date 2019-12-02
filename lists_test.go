package unisender_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alexeyco/unisender"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lists request", func() {
	Context("createList", func() {
		var (
			apiKey   string
			language string
			id       int64
			title    string
		)

		BeforeEach(func() {
			apiKey = randomAPIKey()
			language = randomLanguage()
			id = randomInt64(999, 99999)
			title = randomString(64)
		})

		It("should return correct data", func() {
			expectedUrl := "https://api.unisender.com/" + language + "/api/getLists"

			client := &http.Client{
				Transport: RoundTripFunc(func(req *http.Request) *http.Response {
					Expect(req.URL.String()).To(Equal(expectedUrl), "Incorrect request URL")
					Expect(req.Method).To(Equal(http.MethodPost), "Incorrect request method")

					body := fmt.Sprintf(`{"result":[{"id":%d,"title":%s}]}`, id, title)

					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
						Header:     make(http.Header),
					}
				}),
			}

			usndr := unisender.New(apiKey)
			usndr.SetLanguage(language)
			usndr.SetClient(client)

			lists, err := usndr.GetLists()

			Expect(len(lists)).To(Equal(1), "Wrong lists length returned")
			Expect(err).To(BeNil(), "Error should be nil")
		})
	})
})
