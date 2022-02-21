package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kennykarnama/citcall-test/api"
	"github.com/kennykarnama/citcall-test/client/citcall"

	"github.com/alexflint/go-arg"
)

var args struct {
	Port                   string        `arg:"-p,--port" default:"8080" help:"port for serving http server"`
	ExternalRequestTimeout time.Duration `arg:"--external-request-timeout" default:"15s" help:"http request timeout for external client"`
	CitcallBaseURL         string        `arg:"--citcall-base-url" default:"https://citcall.com" help:"citcall base url"`
}

func main() {
	arg.MustParse(&args)

	// http client
	httpCli := &http.Client{}
	citcallClient := citcall.NewHttpClient(httpCli, args.CitcallBaseURL)
	apiHandler := api.NewHttpHandler(citcallClient, args.ExternalRequestTimeout)

	http.HandleFunc("/countries", apiHandler.GetCountries)

	log.Printf("Serving HTTP on Port: %s", args.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", args.Port), nil)
}
