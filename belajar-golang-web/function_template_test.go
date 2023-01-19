package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name Is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Budi" }}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Bobi",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Print(string(body))
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ len "Budi" }}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Bobi",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Print(string(body))
}

func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {

	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ upper .Name }}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Bobi",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Print(string(body))
}

func TemplateFunctionCreateGlobalPipelines(writer http.ResponseWriter, request *http.Request) {

	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},

		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Bobi",
	})
}

func TestTemplateFunctionCreateGlobalPipelines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobalPipelines(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
