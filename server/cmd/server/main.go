package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"fani/api"
	faniv1connect "fani/gen/fani/v1/v1connect"
	"fani/internal"

	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var settingsPath string

func main() {
	flag.StringVar(&settingsPath, "c", "settings.toml", "path to settings.toml")
	flag.Parse()

	settings, err := api.LoadSettings(settingsPath)
	if err != nil {
		panic(err)
	}

	faniServer := &internal.FaniServer{
		Settings: settings,
	}
	go faniServer.Start(context.Background())

	mux := http.NewServeMux()
	// sets up connect-rpc paths.
	mux.Handle(faniv1connect.NewFaniServiceHandler(faniServer))
	// sets up static file paths for hosting the web server.
	mux.Handle("/", http.FileServer(http.Dir("static")))

	address := "localhost:8080"
	fmt.Printf("serving %s\n", address)
	http.ListenAndServe(address, withCORS(h2c.NewHandler(mux, &http2.Server{})))
}

func withCORS(connectHandler http.Handler) http.Handler {
	return cors.
		New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: connectcors.AllowedMethods(),
			AllowedHeaders: connectcors.AllowedHeaders(),
			ExposedHeaders: connectcors.ExposedHeaders(),
			MaxAge:         7200,
		}).
		Handler(connectHandler)
}
