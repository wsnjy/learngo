package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ResponseCode(write http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		write.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(write, "name is empty")
	} else {
		fmt.Fprintf(write, "Hello %s", name)
	}
}

func TestResponseCodeSuccess(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Bobi", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	assert.Equal(t, response.StatusCode, 200)
	assert.Equal(t, bodyString, "Hello Bobi")
}

func TestResponseCodeFailed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	assert.Equal(t, response.StatusCode, 400)
	assert.Equal(t, bodyString, "name is empty")
}
