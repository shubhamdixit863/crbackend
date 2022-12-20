package app

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"microservicesgo/mocks"
	"net/http/httptest"
	"testing"
)

type TestResponseByID struct {
	Data string `json:"data"`
}

// Integration tests for handlers
func Test_HealthHandler(t *testing.T) {

	req := httptest.NewRequest("GET", "http://localhost:8080/health", nil)

	w := httptest.NewRecorder()
	h := UserHandlers{
		service:    &mocks.MockListingService{},
		fileUpload: &mocks.MockFileUploadService{},
	}
	h.HealthCheck(w, req)

	res := w.Result()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	responseByID := TestResponseByID{}
	err = json.Unmarshal(data, &responseByID)

	assert.Equal(t, "Server Working Properly", responseByID.Data, "")
	assert.Equal(t, 200, w.Code, "Status code")

}
