package routes

import (
	"net/http"
	"wiedzze_server_go/src/controllers"

	"go.mongodb.org/mongo-driver/mongo"
)

// UserRouter funkcja inicjalizująca routing dla użytkowników.
func UserRouter(client *mongo.Client) {
	// Rejestrujemy handler CreateUser dla ścieżki "/users/create".
	// Przekazujemy klienta bazy danych do funkcji CreateUser, aby mogła ona używać tego klienta wewnątrz handlera.
	http.HandleFunc("/users/create", controllers.CreateUser(client))
}