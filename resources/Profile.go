package resources

import (
	"github.com/magiconair/properties"
	"log"
	"os"
	"path/filepath"
)

func GetApplicationProfile() *properties.Properties {

	if len(os.Args) > 1 && os.Args[1] == "dev" {

		log.Println("APPLICATION STARTING IN DEV PROFILE")
		return properties.MustLoadFile(os.Getenv("CLIQUE_CONFIG")+"/clique-mobile-gateway.properties", properties.UTF8)
	} else if len(os.Args) > 1 && os.Args[1] == "prod" {

		panic("Production properties not configured")
	} else {

		log.Println("APPLICATION STARTING IN LOCAL PROFILE")

		wd, _ := os.Getwd()
		propertyPath := filepath.Join([]string{wd, "resources", "application.properties"}...)

		return properties.MustLoadFile(propertyPath, properties.UTF8)
	}
}
