package controllers_test

import (
	"appengine/aetest"
	"github.com/drborges/wetagit/api"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterDevice(t *testing.T) {
	aeinstance, _ := aetest.NewInstance(nil)
	defer aeinstance.Close()

	request, _ := aeinstance.NewRequest("POST", "/tags", strings.NewReader(`{
		"value:" : "golang"
	}`))

	record := httptest.NewRecorder()
	api.Router().ServeHTTP(record, request)

	if record.Code != http.StatusCreated {
		t.Errorf("Expected %v, got %v", http.StatusCreated, record.Code)
	}
}
