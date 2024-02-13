package config

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ładujemy zmienne
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		panic("Nie udało się znaleźć plików environmentu")
	}
}
// Pobieramy zmienną środowiskową MomgoURI z .env
func GetMongoURI() string {
	return os.Getenv("MongoURI")
}

func InitMongoClient(mongoURI string) (*mongo.Client, error) {
	// ustawienia klienta z URI
	clientOptions := options.Client().ApplyURI(mongoURI)
	// utworzenie kontekstu z timeoutem
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	// następnie łączymy się z mongodb
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// pobieramy zzmienną srodowiskową PORT z .env
func GetPort() string {
	port:=os.Getenv("PORT")
	if port == "" {
		return "7070"
	}
	return port
}

