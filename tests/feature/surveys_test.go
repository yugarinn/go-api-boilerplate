package tests

import (
	"testing"
	"net/http/httptest"
)

func TestGetSurveys(t *testing.T) {
	request := httptest.NewRequest("GET", "/surveys", nil)
	response := httptest.NewRecorder()
}
