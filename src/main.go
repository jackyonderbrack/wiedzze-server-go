package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"wiedzze_server_go/src/config"
)

type healthcheckResponse struct {
	Message string `json:"message"`
}

func main() {
	// ładujemy zmienne środowiskowe
	log.Println("Ładowanie zmiennych")
	config.LoadEnv()
	
	// pobieramy MongoURI
	log.Println("Ładowanie URI")
	mongoURI := config.GetMongoURI()

	// Inicjalizujemy klienta mongoDb
	log.Println("Inicjacja klienta mongoDB za pomocą mongoURI")
	databaseClient, err := config.InitMongoClient(mongoURI)
	if err != nil {
		log.Fatalf("Nie udało się zainicjować klienta MongoDB: %v", err)
	} else {
		log.Println("Połączono z bazą danych")
	}

	defer func() {
		if err = databaseClient.Disconnect(context.Background()); err != nil {
			log.Fatalf("Nie udało się poprawnie zamknąć połączenia z bazą danych: %v", err)
		}
	}()

	// od teraz można używać zmiennej 'client' do interakcji z bazą danych

	// pobieramy port z config.go
	port := config.GetPort()
	// Healthckeck
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		response := healthcheckResponse{Message: "pong"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Uwaga: Użyj log.Fatal zamiast log.Panicf do uruchomienia serwera
	log.Printf("Uruchamianie serwera na porcie %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Błąd podczas uruchamiania serwera: ", err)
	}
}