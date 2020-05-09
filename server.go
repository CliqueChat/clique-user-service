package main

import (
	"github.com/CliqueChat/clique-common-lib/helpers"
	"github.com/magiconair/properties"
	"log"
	"net/http"
	"os"
)

var prop = properties.MustLoadFile(os.Getenv("CLIQUE_CONFIG")+"/clique-user-service.properties", properties.UTF8)

func main() {

	// Read host and port from property file
	host, _ := prop.Get(helpers.HOST)
	port, _ := prop.Get(helpers.PORT)

	log.Println("STARTING APPLICATION IN " + host + ":" + port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
