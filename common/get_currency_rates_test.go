package common_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/common"
	"github.com/alexeyco/unisender/test"
)

func TestGetCurrencyRatesRequest_Execute(t *testing.T) {
	expectedResult := []common.CurrencyRate{
		{
			ID:            test.RandomInt64(9999, 999999),
			Code:          test.RandomString(12, 36),
			RateToUSD:     test.RandomFloat64(),
			MinPaymentSum: test.RandomFloat64(),
			IsVisible:     true,
		},
	}

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := common.GetCurrencyRates(req).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatalf("Results should be equal")
	}
}
