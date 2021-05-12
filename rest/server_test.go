// I opted for keeping server_test.go in the same package in order to make test internal functions from server.go
package restserver

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ExpectedTestCase struct {
	X     *float32
	Y     *float32
	Error error
}
type RequestTestCase struct {
	Request  *http.Request
	Expected ExpectedTestCase
}

func TestExtractValues(t *testing.T) {
	r1, _ := http.NewRequest(http.MethodGet, "/multiply?x=2&y=4", nil)
	r2, _ := http.NewRequest(http.MethodGet, "/multiply?&y=4", nil)
	r3, _ := http.NewRequest(http.MethodGet, "/multiply?&x=4", nil)
	r4, _ := http.NewRequest(http.MethodGet, "/multiply?x=4,23&y=1.3", nil)
	r5, _ := http.NewRequest(http.MethodGet, "/multiply?x=4.23&y=1,,3", nil)
	r6, _ := http.NewRequest(http.MethodGet, "/multiply?x=1.5&y=1.5", nil)
	r7, _ := http.NewRequest(http.MethodGet, "/multiply?x=-1.5&y=1.5", nil)
	var x1 float32 = 2
	var y1 float32 = 4
	var x6 float32 = 1.5
	var y6 float32 = 1.5
	var x7 float32 = -1.5
	var y7 float32 = 1.5
	testCases := []RequestTestCase{
		{
			Request: r1,
			Expected: ExpectedTestCase{
				X:     &x1,
				Y:     &y1,
				Error: nil,
			},
		},
		{
			Request: r2,
			Expected: ExpectedTestCase{
				X:     nil,
				Y:     nil,
				Error: errors.New("incorrect number params"),
			},
		},
		{
			Request: r3,
			Expected: ExpectedTestCase{
				X:     nil,
				Y:     nil,
				Error: errors.New("incorrect number params"),
			},
		},
		{
			Request: r4,
			Expected: ExpectedTestCase{
				X:     nil,
				Y:     nil,
				Error: errors.New(`strconv.ParseFloat: parsing "4,23": invalid syntax`),
			},
		},
		{
			Request: r5,
			Expected: ExpectedTestCase{
				X:     nil,
				Y:     nil,
				Error: errors.New(`strconv.ParseFloat: parsing "1,,3": invalid syntax`),
			},
		},
		{
			Request: r6,
			Expected: ExpectedTestCase{
				X:     &x6,
				Y:     &y6,
				Error: nil,
			},
		},
		{
			Request: r7,
			Expected: ExpectedTestCase{
				X:     &x7,
				Y:     &y7,
				Error: nil,
			},
		},
	}

	for _, tcase := range testCases {
		x, y, err := extractValues(tcase.Request)
		if err != nil && tcase.Expected.Error != nil {
			assert.Equal(t, err.Error(), tcase.Expected.Error.Error(), "")
		} else {
			assert.Equal(t, err, tcase.Expected.Error, "")
		}
		if x != nil && tcase.Expected.X != nil {
			assert.Equal(t, *x, *tcase.Expected.X, "")
		} else {
			assert.Equal(t, x, tcase.Expected.X, "")
		}
		if y != nil && tcase.Expected.Y != nil {
			assert.Equal(t, *y, *tcase.Expected.Y, "")
		} else {
			assert.Equal(t, y, tcase.Expected.Y, "")
		}
	}

}

func TestMultiplyHandlerV1(t *testing.T) {
	// fail case
	req, _ := http.NewRequest("GET", "/multiply?x=2", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MultiplyRestHandlerV1)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusBadRequest, "")
	byteBody, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("error at reading body response%v\n", err)
	}
	errMsg := "Incorrect parameters setted. Error: incorrect number params"
	expectedPayload := Response{
		Data:  nil,
		Error: &errMsg,
	}
	var actualPayload Response
	_ = json.Unmarshal(byteBody, &actualPayload)
	t.Logf("%v\n", actualPayload)
	assert.Equal(t, expectedPayload, actualPayload, "")

	// success case
	req, _ = http.NewRequest("GET", "/multiply?x=2&y=4", nil)
	expectedData := float32(8.0)
	expectedSuccessBody := Response{
		Data:  &expectedData,
		Error: nil,
	}
	handler = http.HandlerFunc(MultiplyRestHandlerV1)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK, "")
	byteBody, err = ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("error at reading body response%v\n", err)
	}
	_ = json.Unmarshal(byteBody, &actualPayload)
	assert.Equal(t, actualPayload, expectedSuccessBody, "")

	req, _ = http.NewRequest("GET", "/multiply?x=2e20&y=4e20", nil)
	handler = http.HandlerFunc(MultiplyRestHandlerV1)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusInternalServerError, "")
	byteBody, err = ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("error at reading body response%v\n", err)
	}
	_ = json.Unmarshal(byteBody, &actualPayload)
	errMsg = "Number is really big and exceeded maximum range value"
	expectedPayload = Response{
		Data:  nil,
		Error: &errMsg,
	}
	assert.Equal(t, actualPayload, expectedPayload, "")
}
