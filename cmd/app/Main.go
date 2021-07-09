package main

import (
	"CODE/config"
	server "CODE/server/app"
	"CODE/server/router"
	lr "CODE/util/logger"
	"fmt"
	"log"
	"net/http"
)

func main() {
	appConf := config.AppConfig()
	logger := lr.New(appConf.Server.Debug)
	application := server.New(logger)
	appRouter := router.New(application)

	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)

	address := fmt.Sprintf(":%d", appConf.Server.Port)
	logger.Info().Msgf("Starting server %v", address)

	log.Printf("Starting server %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
