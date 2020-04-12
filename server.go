package main

import (
	"github.com/CliqueChat/clique-common-lib/helpers"
	"github.com/magiconair/properties"
	"log"
	"net/http"
)

var prop = properties.MustLoadFile("./properties/config.properties", properties.UTF8)

func main() {

	// Read host and port from property file
	host , _ := prop.Get(helpers.HOST)
	port , _ := prop.Get(helpers.PORT)
}