package generate

import (
	"github.com/wzshiming/go-swagger/swagger"
	"github.com/wzshiming/go-swagger/swaggergen"
	"github.com/wzshiming/go-swagger/utils"
	"github.com/wzshiming/pot"
)

func GenerateDocs(routers, controllers, swagge string) error {
	s := &swagger.Swagger{}
	swaggergen.GB(s, routers, controllers)
	err := utils.WriteFile(s, swagge)
	if err != nil {
		return err
	}
	pot.Println("Generate docs")
	return nil
}