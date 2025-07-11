package main

import (
	"filter-service/configs"
	"filter-service/internal/raw"
	"filter-service/metrics"
	"filter-service/middleware"
	"filter-service/pkg/db"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)

	// Prometheus
	go func() {
		metrics.Listen("0.0.0.0:8082")
	}()

	router := http.NewServeMux()

	//Repositories
	rawRepository := raw.NewRawRepository(db)

	//Services
	rawService := raw.NewRawService(rawRepository)

	//Handlers
	raw.NewRawHandler(router, raw.RawHandlerDeps{
		RawService: rawService,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.RawLog,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}
	fmt.Println("Server starting in port 8081")
	server.ListenAndServe()

}
