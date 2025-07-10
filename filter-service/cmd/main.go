package main

import (
	"filter-service/configs"
	"filter-service/internal/raw"
	"filter-service/middleware"
	"filter-service/pkg/db"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)

	router := http.NewServeMux()

	//Repositories
	rawRepository := raw.NewRawRepository(db)

	//Services
	rawService := raw.NewRawService(rawRepository)

	//Handlers
	raw.NewRawHandler(router, raw.RawHandlerDeps{
		RawService: rawService,
	})

	//Middlewares
	stack := middleware.Chain(
		middleware.Logging,
	)
	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}
	fmt.Println("Server starting in port 8081")
	server.ListenAndServe()

}

// func main() {
// 	conf := configs.LoadConfig()
// 	log.Printf("Db %s | Token %s", conf.Db.Dsn, conf.Auth.Secret)

// 	db := db.NewDb(conf)
// 	go metrics.StartPrometheus() // Prometheus

// 	router := http.NewServeMux()

// 	//Repositories
// 	linkRepository := link.NewLinkRepository(db)
// 	userRepository := user.NewUserRepository(db)

// 	//Servises
// 	authServise := auth.NewAuthService(userRepository)

// 	// Handlers
// 	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
// 		Config:      conf,
// 		AuthService: authServise,
// 	})
// 	link.NewLinkHandler(router, link.LinkHandlerDeps{
// 		LinkRepository: linkRepository,
// 	})

// 	//Middlewares
// 	stack := middleware.Chain(
// 		middleware.CORS,
// 		middleware.Logging,
// 	)

// 	server := http.Server{
// 		Addr:    ":8081",
// 		Handler: stack(router),
// 	}
// 	fmt.Println("Server starting ...")
// 	server.ListenAndServe()
// }
