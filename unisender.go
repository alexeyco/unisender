package unisender

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiEndpointPattern = "https://api.unisender.com/%s/api/%s"

// UniSender api client
type UniSender struct {
	language string
	apiKey   string
	logger   Logger
	client   *http.Client
}

// SetLanguage sets API response language
func (u *UniSender) SetLanguage(language string) {
	u.language = language
}

func (u *UniSender) SetLogger(logger Logger) {
	u.logger = logger
}

// SetClient sets custom http.Client to UniSender client
func (u *UniSender) SetClient(client *http.Client) {
	u.client = client
}

// request makes request to UniSender API and parses response
func (u *UniSender) request(method string, v interface{}, options ...Option) (err error) {
	uri := fmt.Sprintf(apiEndpointPattern, u.language, method)

	data := url.Values{}
	data.Add("api_key", u.apiKey)

	for _, o := range options {
		data.Add(o.name, o.value)
	}

	u.logRequest(uri, data)

	var resp *http.Response
	if resp, err = u.client.PostForm(uri, data); err != nil {
		return
	}

	u.logResponseBody(resp)
	if v == nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		return ErrWrongStatusCode
	}

	response := Response{
		Result: v,
	}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return
	}

	if err = response.Err(); err != nil {
		return
	}

	v = response.Result

	return
}

func (u *UniSender) logRequest(uri string, data url.Values) {
	if u.logger == nil {
		return
	}

	params := map[string]interface{}{}
	for p, v := range data {
		params[p] = v[0]
	}

	u.logger.Println(fmt.Sprintf("POST %s", uri), params)
}

func (u *UniSender) logResponseBody(resp *http.Response) {
	if u.logger == nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	u.logger.Println("Response", map[string]interface{}{
		"statusCode": resp.StatusCode,
		"body":       string(body),
	})
}

// New returns new UniSender client
func New(apiKey string) *UniSender {
	return &UniSender{
		language: DefaultLanguage,
		apiKey:   apiKey,
		client:   http.DefaultClient,
	}
}
