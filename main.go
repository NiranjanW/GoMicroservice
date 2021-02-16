package main

import (
	"fmt"
	"log"
	"myMicroService/homepage"
	"myMicroService/server"
	"net/http"
	"os"
)

var (
	GcukCertFile    = os.Getenv("GCUK_CERT_FILE")
	GcukKeyFile     = os.Getenv("GCUK_KEY_FILE")
	GcukServiceAddr = os.Getenv("GCUK_SERVICE_ADDR")
)



func main() {
	logger := log.New(os.Stdout, "gcul", log.LstdFlags|log.Lshortfile)
	h := homepage.NewHandlers(logger)

	fmt.Println("Hello World")
	mux := http.NewServeMux()
	//mux.HandleFunc("/", h.HomeHandler)
	h.SetUpRoutes(mux)
	srv := server.New(mux, GcukServiceAddr)

	err := srv.ListenAndServeTLS(GcukCertFile, GcukKeyFile)
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}

}
