package routes

import (
	"fmt"
	"net/http"

	"github.com/nicolasrsaraiva/controllers"
)

func TrainerHandleRequests() {
	fmt.Println("Rotas de trainer iniciadas")
	http.HandleFunc("/trainer", controllers.CreateTrainer)
	http.HandleFunc("/trainer/", controllers.GetTrainer)
	http.HandleFunc("/trainers", controllers.GetTrainers)
	http.HandleFunc("/trainer/create-team", controllers.CreateTeam)
}

func PokemonHandleRequests() {
	fmt.Println("Rotas de pokemon iniciadas")
	http.HandleFunc("/pokemon/name/", controllers.GetPokemonByName)
	http.HandleFunc("/pokemon/id/", controllers.GetPokemonById)
}
