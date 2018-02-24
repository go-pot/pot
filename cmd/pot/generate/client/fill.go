package client

import (
	"strings"

	"github.com/wzshiming/go-swagger/swagger"
)

type FilledPackage struct {
	Paths       []*FilledPath
	Definitions []*FilledDefinitions
}

func (f *FilledPackage) Parse(swa *swagger.Swagger) {
	for k, v := range swa.Paths {
		if v.Get != nil {
			pp := &FilledPath{Path: k, Method: "Get"}
			pp.Parse(v.Get)
			f.Paths = append(f.Paths, pp)
		}

		if v.Put != nil {
			pp := &FilledPath{Path: k, Method: "Put"}
			pp.Parse(v.Put)
			f.Paths = append(f.Paths, pp)
		}

		if v.Post != nil {
			pp := &FilledPath{Path: k, Method: "Post"}
			pp.Parse(v.Post)
			f.Paths = append(f.Paths, pp)
		}

		if v.Delete != nil {
			pp := &FilledPath{Path: k, Method: "Delete"}
			pp.Parse(v.Delete)
			f.Paths = append(f.Paths, pp)
		}

		if v.Options != nil {
			pp := &FilledPath{Path: k, Method: "Options"}
			pp.Parse(v.Options)
			f.Paths = append(f.Paths, pp)
		}

		if v.Head != nil {
			pp := &FilledPath{Path: k, Method: "Head"}
			pp.Parse(v.Head)
			f.Paths = append(f.Paths, pp)
		}

		if v.Patch != nil {
			pp := &FilledPath{Path: k, Method: "Patch"}
			pp.Parse(v.Patch)
			f.Paths = append(f.Paths, pp)
		}
	}

	for k, v := range swa.Definitions {
		pp := &FilledDefinitions{Name: k}
		pp.Parse(&v)
		f.Definitions = append(f.Definitions, pp)
	}
}

type FilledProperties struct {
	Name        string
	Type        string
	Description string
}

type FilledDefinitions struct {
	Name        string
	Properties  []*FilledProperties
	Description string
}

func (f *FilledDefinitions) Parse(swa *swagger.Schema) {
	f.Description = swa.Description
	for k, v := range swa.Properties {
		f.Properties = append(f.Properties, &FilledProperties{
			Name:        k,
			Type:        ToGoType(v.Format, v.Type),
			Description: v.Description,
		})
	}
}

type FilledPath struct {
	Func               string
	Method             string
	Path               string
	ParametersHeader   []*FilledParameters
	ParametersBody     []*FilledParameters
	ParametersQuery    []*FilledParameters
	ParametersFormData []*FilledParameters
	ParametersPath     []*FilledParameters
	Responses          []*FilledResponses
}

func (f *FilledPath) Parse(swa *swagger.Operation) {
	f.Func = swa.OperationID
	for _, v := range swa.Parameters {

		switch v.In {
		case "header":
			f.ParametersHeader = append(f.ParametersHeader, &FilledParameters{
				Name:        v.Name,
				Type:        ToGoType(v.Format, v.Type),
				Description: v.Description,
			})
		case "body":

			if v.Schema != nil {
				ff := &FilledParameters{
					Name:        v.Name,
					Description: v.Description,
				}
				if v.Schema.Type == "array" {
					ff.Type = "[]*" + strings.TrimPrefix(v.Schema.Items.Ref, "#/definitions/")
				} else {
					ff.Type = "*" + strings.TrimPrefix(v.Schema.Ref, "#/definitions/")
				}
				f.ParametersBody = append(f.ParametersBody, ff)
			}

		case "query":
			f.ParametersQuery = append(f.ParametersQuery, &FilledParameters{
				Name:        v.Name,
				Type:        ToGoType(v.Format, v.Type),
				Description: v.Description,
			})
		case "formData":
			f.ParametersFormData = append(f.ParametersFormData, &FilledParameters{
				Name:        v.Name,
				Type:        ToGoType(v.Format, v.Type),
				Description: v.Description,
			})
		case "path":
			f.ParametersPath = append(f.ParametersPath, &FilledParameters{
				Name:        v.Name,
				Type:        ToGoType(v.Format, v.Type),
				Description: v.Description,
			})
		}
	}

	for k, v := range swa.Responses {
		if v.Schema != nil {
			ff := &FilledResponses{
				Code:        k,
				Unmarshal:   true,
				Description: v.Description,
			}
			if v.Schema.Type == "array" {
				ff.Type = "[]*" + strings.TrimPrefix(v.Schema.Items.Ref, "#/definitions/")
			} else {
				ff.Type = "*" + strings.TrimPrefix(v.Schema.Ref, "#/definitions/")
			}
			f.Responses = append(f.Responses, ff)
		} else {
			f.Responses = append(f.Responses, &FilledResponses{
				Code:        k,
				Type:        "[]byte",
				Unmarshal:   false,
				Description: v.Description,
			})
		}
	}
}

type FilledParameters struct {
	Name        string
	Type        string
	Description string
}

type FilledResponses struct {
	Code        string
	Type        string
	Unmarshal   bool
	Description string
}

func ToGoType(st ...string) string {
	for _, v := range st {
		if len(v) != 0 {
			d := toGoType[v]
			if len(d) != 0 {
				return d
			}
		}
	}
	return toGoType[""]
}

var toGoType = map[string]string{
	"":          "string",
	"boolean":   "bool",
	"int64":     "int",
	"int32":     "int",
	"int":       "int",
	"integer":   "int",
	"string":    "string",
	"number":    "float64",
	"double":    "float64",
	"float":     "float64",
	"date-time": "time.Time",
}
