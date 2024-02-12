package main

import (
	"fmt"
	"log"
	"net/http"
	"wiedzze_server_go/src/routes"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Witaj w serwerze WiedzŻe")
	})

	routes.UserRouter()

	fmt.Println("Serwer działa na porcie 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	
}