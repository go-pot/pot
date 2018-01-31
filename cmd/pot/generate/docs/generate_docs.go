package docs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/wzshiming/go-swagger/swagger"
	"github.com/wzshiming/go-swagger/swaggergen"
	yaml "gopkg.in/yaml.v2"
)

func GenerateDocs(routers, controllers, swagge string) error {
	s := &swagger.Swagger{}
	swaggergen.GB(s, routers, controllers)
	err := WriteFile(s, swagge)
	if err != nil {
		return err
	}
	return nil
}

func WriteFile(rootapi interface{}, basepath string) error {
	ext := filepath.Ext(basepath)
	if ext == "" || ext == ".json" {
		dt, err := json.MarshalIndent(rootapi, "", "    ")
		if err != nil {
			return err
		}
		jf := filepath.Join(basepath, "swagger.json")
		dt = append(dt, '\n')

		d, _ := ioutil.ReadFile(jf)
		if string(d) == string(dt) {
			fmt.Println("Unchanged docs swagger.json")
		} else {
			err = ioutil.WriteFile(jf, dt, 0666)
			if err != nil {
				return err
			}
			fmt.Println("Generate docs swagger.json")
		}
	}

	if ext == "" || ext == ".yml" {
		dt, err := yaml.Marshal(rootapi)
		if err != nil {
			return err
		}
		yf := filepath.Join(basepath, "swagger.yml")

		d, _ := ioutil.ReadFile(yf)
		if string(d) == string(dt) {
			fmt.Println("Unchanged docs swagger.yml")
		} else {
			err = ioutil.WriteFile(yf, dt, 0666)
			if err != nil {
				return err
			}
			fmt.Println("Generate docs swagger.yml")
		}
	}
	return nil
}
