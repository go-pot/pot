package generate

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"text/template"

	"github.com/wzshiming/go-swagger/swagger"
	"github.com/wzshiming/go-swagger/swaggergen"
	"github.com/wzshiming/pot"
)

func GenerateRouter(routers string, controllers, out string) error {

	swagger := &swagger.Swagger{}
	swaggergen.GB(swagger, routers, controllers)

	t := "temp"
	temp := template.New(t)
	_, err := temp.Parse(mkRouter)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)
	err = temp.ExecuteTemplate(buf, "temp", swagger)
	if err != nil {
		return err
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(out, src, 0666)
	if err != nil {
		return err
	}

	pot.Println("Generate routers")
	return nil
}

var mkRouter = `// Code generated by "pot gen rou"; DO NOT EDIT.
package routers

import (
	controllers "{{.Extensions.Package}}"
	"net/http"

	"github.com/wzshiming/pot"
	"github.com/wzshiming/pot/router"
)

func init() {
	p := pot.Default
	r := router.NewRouter()

	paths := r.PathPrefix("{{.BasePath}}").Subrouter()

	// init controllers
	{
		{{range $k, $v := .Paths}}
		{{with .Get}}
		paths.Methods(http.MethodGet).Path("{{.OperationID}}").HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				t := controllers.{{.Extensions.Controllers}}{}
				t.Init(p, w, r)
				t.{{.Extensions.Methods}}()
			})
		{{end}}
		{{with .Post}}
		paths.Methods(http.MethodPost).Path("{{.OperationID}}").HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				t := controllers.{{.Extensions.Controllers}}{}
				t.Init(p, w, r)
				t.{{.Extensions.Methods}}()
			})
		{{end}}
		{{with .Put}}
		paths.Methods(http.MethodPut).Path("{{.OperationID}}").HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				t := controllers.{{.Extensions.Controllers}}{}
				t.Init(p, w, r)
				t.{{.Extensions.Methods}}()
			})
		{{end}}
		{{with .Delete}}
		paths.Methods(http.MethodDelete).Path("{{.OperationID}}").HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				t := controllers.{{.Extensions.Controllers}}{}
				t.Init(p, w, r)
				t.{{.Extensions.Methods}}()
			})
		{{end}}
		{{with .Options}}
		paths.Methods(http.MethodOptions).Path("{{.OperationID}}").HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				t := controllers.{{.Extensions.Controllers}}{}
				t.Init(p, w, r)
				t.{{.Extensions.Methods}}()
			})
		{{end}}
		{{with .Head}}
		paths.Methods(http.MethodHead).Path("{{.OperationID}}").HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				t := controllers.{{.Extensions.Controllers}}{}
				t.Init(p, w, r)
				t.{{.Extensions.Methods}}()
			})
		{{end}}
		{{with .Patch}}
		paths.Methods(http.MethodPatch).Path("{{.OperationID}}").HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				t := controllers.{{.Extensions.Controllers}}{}
				t.Init(p, w, r)
				t.{{.Extensions.Methods}}()
			})
		{{end}}
		{{end}}
	}

	r.Walk(router.PrintRouter)
	p.UseHandler(r)
}
`