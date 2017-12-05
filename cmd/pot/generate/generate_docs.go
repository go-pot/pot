package generate

import (
	"github.com/wzshiming/go-swagger/swagger"
	"github.com/wzshiming/go-swagger/swaggergen"
	"github.com/wzshiming/go-swagger/utils"
	"github.com/wzshiming/pot"
)

func GenerateDocs(routers, swagge string) {
	s := &swagger.Swagger{}
	swaggergen.GB(s, routers)
	utils.WriteFile(s, swagge)
	pot.Println("Generate docs")
}
