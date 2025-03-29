package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"terraform-provider-pokemon/internal/provider"
)

const version string = "dev"

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "enable support for debuggers")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "github.com/fionn/terraform-provider-pokemon",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
