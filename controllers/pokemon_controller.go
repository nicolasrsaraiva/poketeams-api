package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nicolasrsaraiva/database"
	"github.com/nicolasrsaraiva/models"
)

func GetPokemonByName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	name := r.URL.Query().Get("name")

	db := database.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM pokemons WHERE name=$1")

	if err != nil {
		http.Error(w, "Error in pokemon name", http.StatusBadRequest)
	}
	defer stmt.Close()

	var pokemon models.Pokemon

	row := stmt.QueryRow(name)

	err = row.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Hp, &pokemon.Def, &pokemon.Defm, &pokemon.Atk, &pokemon.Spatk, &pokemon.Speed)

	if err != nil {
		http.Error(w, "Error in query row", http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(pokemon)
	if err != nil {
		http.Error(w, "Error encoding pokemon", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
}

func GetPokemonById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	id := r.URL.Query().Get("id")
	idPkm, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Error in pokemon id", http.StatusBadRequest)
	}

	db := database.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM pokemons WHERE ID=$1")
	if err != nil {
		http.Error(w, "Error in query select", http.StatusInternalServerError)
	}
	defer stmt.Close()

	var pokemon models.Pokemon

	row := stmt.QueryRow(idPkm)

	err = row.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Hp, &pokemon.Def, &pokemon.Defm, &pokemon.Atk, &pokemon.Spatk, &pokemon.Speed)

	if err != nil {
		http.Error(w, "Error in pokemon scan", http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(pokemon)
	if err != nil {
		http.Error(w, "Error in pokemon encoding", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}
