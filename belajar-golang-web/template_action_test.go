package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateIf(write http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(write, "if.gohtml", map[string]interface{}{
		"Name": "Bobi",
	},
	)
}

func TestTemplateIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateIf(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Print(string(body))
}

func TemplateActionOperator(write http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(write, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Action Operator",
		"FinalValue": 50,
	},
	)
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Print(string(body))
}

func TemplateActionRange(write http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(write, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Operator",
		"Hobbies": []string{
			"Game", "Read", "Code",
		},
	},
	)
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Print(string(body))
}

func TemplateActionWith(write http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	t.ExecuteTemplate(write, "address.gohtml", map[string]interface{}{
		"Title": "Template Action Operator",
		"Name":  "Bobi Bikul",
		"Address": map[string]interface{}{
			"Street": "Jimbaran",
			"City":   "Bali",
		},
	},
	)
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Print(string(body))
}
