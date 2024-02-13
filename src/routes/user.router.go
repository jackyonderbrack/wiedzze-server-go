package routes

import "net/http"

func UserRouter() {
	http.HandleFunc("/users", usersHandler)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	
}