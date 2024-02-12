package routes

import "net/http"

func UserRouter() {
	http.HandleFunc("/users", usersHandler)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/users", controllers.GetUser)
	http.HandleFunc("/users/create", controllers.CreateUser)
	http.HandleFunc("/users/update", controllers.UpdateUser)
	http.HandleFunc("/users/delete", controllers.DeleteUser)
}