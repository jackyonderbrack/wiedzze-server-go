package main

import (
	"context"
	"log"
	"wiedzze_server_go/src/config"
)

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
}