package main

import (
	"flag"
	"log"

	gen "github.com/s-you/openapi-error-code-gen"
)

var (
	folderName  = flag.String("path", "apierrors", "path to output error codes")
	pkgName     = flag.String("pkg", "apierrors", "package name")
)

func main() {
	flag.Parse()

	// swagger, err := openapi.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	if err := gen.Generate(swagger, *pkgName, *folderName); err != nil {
		log.Fatal(err)
	}
}

