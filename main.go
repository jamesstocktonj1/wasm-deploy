package main

import (
	"fmt"
	"net/http"

	wasihttp "github.com/jamesstocktonj1/componentize-sdk/p3/exports/http"
)

func init() {
	wasihttp.HandleFunc(handler)
}

const (
	FunctionVersion = "0.1.0" // Function Version
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Function, Version %s!", FunctionVersion)
}

func main() {}
