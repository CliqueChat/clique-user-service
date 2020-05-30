package resources

import (
	"github.com/magiconair/properties"
	"os"
	"path/filepath"
)

func GetApplicationProfile() (*properties.Properties, string) {

	if len(os.Args) > 1 && os.Args[1] == "dev" {

		return properties.MustLoadFile(os.Getenv("CLIQUE_CONFIG")+"/clique-mobile-gateway.properties", properties.UTF8), "DEV"

	} else if len(os.Args) > 1 && os.Args[1] == "prod" {

		panic("Production properties not configured")
	} else {

		wd, _ := os.Getwd()
		propertyPath := filepath.Join([]string{wd, "resources", "application.properties"}...)

		return properties.MustLoadFile(propertyPath, properties.UTF8), "LOCAL"
	}
}
