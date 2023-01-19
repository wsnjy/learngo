package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(write http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/layout.gohtml",
		"./templates/footer.gohtml",
	))
	t.ExecuteTemplate(write, "layout", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Bobi",
	},
	)
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Print(string(body))
}
