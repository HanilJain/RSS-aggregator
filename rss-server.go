package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv" //mod to godotenv
)

func main() {
	fmt.Println("hello world")

	// use to export the varabile from env in the shell otherwise it does not load
	godotenv.Load()

	portstring := os.Getenv("PORT")

	if portstring == "" {
		log.Fatal("PORT is not found in env")
	}

	fmt.Println("PORT:", portstring)

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portstring,
	}
	
	log.Printf("Serve starting on PORT %v", portstring)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
