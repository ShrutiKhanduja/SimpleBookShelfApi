package main

import (
	"fmt"
	"log"
	"net/http"

	"CODE/config"
	"CODE/server/router"
	lr "CODE/util/logger"
)

func main() {
	appConf := config.AppConfig()
	appRouter := router.New()
	logger := lr.New(appConf.Server.Debug)

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
