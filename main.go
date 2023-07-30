package main

import (
	"fmt"
	"net/http"

	"github.com/nicolasrsaraiva/routes"
)

func main() {
	routes.TrainerHandleRequests()
	routes.PokemonHandleRequests()
	fmt.Println("Servidor online")
	http.ListenAndServe("localhost:8000", nil)
}
