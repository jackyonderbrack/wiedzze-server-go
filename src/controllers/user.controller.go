package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"wiedzze_server_go/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(databaseClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var user models.User
		//r.Body reprezentuje ciało żądania HTTP, które jest strumieniem danych. Może zawierać dane w formacie JSON, które chcemy przekonwertować na strukturę Go. 
		// json.NewDecoder tworzy nowy dekoder dla strumienia danych JSON z r.Body.
		//.Decode(&user) używa tego dekodera do odczytania strumienia JSON i zapisania przekonwertowanych danych do zmiennej user, która jest wskaźnikiem na strukturę User. Dekodowanie zmienia surowy JSON w strukturalny format Go, którym możemy łatwo manipulować w kodzie.
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Błąd przy dekodowaniu danych użytkownika"})
			return
		}

		// ustawianie dodatkowych informacji przed zapisem
		user.ID = primitive.NewObjectID()
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		// ewentualne hashowanie

		// zapisujemy do bazy danych
		collection := databaseClient.Database("wiedzze").Collection("users")
		_, err := collection.InsertOne(context.Background(), user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Nie można zapisać użykownika "})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}