package unisender_test

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/alexeyco/unisender"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("createList", func() {
	var (
		language    string
		apiKey      string
		requestUrl  string
		expectedUrl string
		id          int64
		title       string
		lists       []unisender.List
		err         error
	)

	BeforeEach(func() {
		language = randomLanguage()
		apiKey = randomAPIKey()

		expectedUrl = fmt.Sprintf("https://api.unisender.com/%s/api/getLists", language)

		id = randomInt64(999, 99999)
		title = randomString(64)
	})

	When("API responds with non-200 status", func() {
		BeforeEach(func() {
			client := &http.Client{
				Transport: roundTrip{
					Before: func(url *url.URL) {
						requestUrl = url.String()
					},
					StatusCode: func() int {
						return http.StatusForbidden
					},
				},
			}

			usndr := unisender.New(apiKey)
			usndr.SetLanguage(language)
			usndr.SetClient(client)

			lists, err = usndr.GetLists()
		})

		It("should sent to correct URL", func() {
			Expect(requestUrl).To(Equal(expectedUrl), "URLs should be equal")
		})

		It("should return error", func() {
			Expect(err).ToNot(BeNil(), "Error should be not nil")
		})

		It("should return no lists", func() {
			Expect(len(lists)).To(Equal(0), "Lists slice should be empty")
		})
	})

	When("data is correct", func() {
		BeforeEach(func() {
			client := &http.Client{
				Transport: roundTrip{
					Before: func(url *url.URL) {
						requestUrl = url.String()
					},
					Body: func() string {
						return fmt.Sprintf(`{"result":[{"id":%d,"title":"%s"}]}`, id, title)
					},
				},
			}

			usndr := unisender.New(apiKey)
			usndr.SetLanguage(language)
			usndr.SetClient(client)

			lists, err = usndr.GetLists()
		})

		It("should sent to correct URL", func() {
			Expect(requestUrl).To(Equal(expectedUrl), "URLs should be equal")
		})

		It("should return no error", func() {
			Expect(err).To(BeNil(), "Error should be nil")
		})

		Context("result", func() {
			It("should have correct length", func() {
				Expect(len(lists)).To(Equal(1), "Wrong lists length returned")
			})

			It("should return correct data", func() {
				list := lists[0]

				Expect(list.ID).To(Equal(id), "ID's should be equal")
				Expect(list.Title).To(Equal(title), "Titles should be equal")
			})
		})
	})
})
