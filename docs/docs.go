package docs

import (
	"io/ioutil"
	"log"

	"github.com/swaggo/swag"
	"gopkg.in/yaml.v3"
)

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "http://127.0.0.1:8080/swagge",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	// SwaggerTemplate:  docTemplate,
}

func init() {
	yfile, err := ioutil.ReadFile("./docs/swagger.yaml")
	if err != nil {
		log.Fatal(err)
	}
	yaml.Marshal(yfile)
	SwaggerInfo.SwaggerTemplate = string(yfile)
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
