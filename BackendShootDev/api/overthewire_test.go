package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

type OverTheWireRequest struct {
	Data string `json:"data"`
}

func TestHandleOverTheWireSuccess(t *testing.T) {
	s := CreateNewServer()

	api := &ApiConfig{}
	s.MountMiddleware(api)
	s.MakeApiHandlers(api)

	payload := OverTheWireRequest{
		Data: "Hello World!",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/overthewire", bytes.NewReader(body))

	response := executeRequest(req, s)

	checkResponseCode(t, http.StatusOK, response.Code)
	require.Equal(t, "\"Hello World!\"", response.Body.String())

}

func TestHandleOverTheWireEmptyBody(t *testing.T) {
	s := CreateNewServer()

	api := &ApiConfig{}
	s.MountMiddleware(api)
	s.MakeApiHandlers(api)

	payload := OverTheWireRequest{
		Data: "",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/overthewire", bytes.NewReader(body))

	response := executeRequest(req, s)

	checkResponseCode(t, http.StatusNoContent, response.Code)
	require.Equal(t, "{\"error\":\"No body in the Payload\"}", response.Body.String())

}

func TestHandleOverTheWireInvalidData(t *testing.T) {
	s := CreateNewServer()

	api := &ApiConfig{}
	s.MountMiddleware(api)
	s.MakeApiHandlers(api)

	body, _ := json.Marshal("{\"data\": \"Hello World")

	req, _ := http.NewRequest("POST", "/overthewire", bytes.NewReader(body))

	response := executeRequest(req, s)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

}
