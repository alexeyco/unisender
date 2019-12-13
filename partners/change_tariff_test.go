package partners_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/partners"
	"github.com/alexeyco/unisender/test"
)

func TestChangeTariffRequest_Execute(t *testing.T) {
	expectedLogin := test.RandomString(12, 36)
	var givenLogin string

	expectedTariffID := test.RandomInt64(9999, 999999)
	var givenTariffID int64

	expectedResult := &partners.ChangeTariffResult{
		TariffID:     test.RandomInt64(9999, 999999),
		ChargedSum:   0,
		ChargedBonus: 0,
		StartTime:    test.RandomTime(12, 365),
		EndTime:      test.RandomTime(12, 365),
	}

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLogin = req.FormValue("login")
		givenTariffID, _ = strconv.ParseInt(req.FormValue("tariff_id"), 10, 64)

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := partners.ChangeTariff(req, expectedLogin, expectedTariffID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLogin != givenLogin {
		t.Fatalf(`Login should be "%s", "%s" given`, expectedLogin, givenLogin)
	}

	if expectedTariffID != givenTariffID {
		t.Fatalf(`Tariff ID should be %d, %d given`, expectedTariffID, givenTariffID)
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatalf("Results should be equal")
	}
}
