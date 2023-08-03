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
		return
	}

	id := r.URL.Query().Get("id")
	idPkm, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Error in pokemon id", http.StatusBadRequest)
		return
	}

	db := database.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM pokemons WHERE ID=$1")
	if err != nil {
		http.Error(w, "Error in query select", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var pokemon models.Pokemon

	row := stmt.QueryRow(idPkm)

	err = row.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Hp, &pokemon.Def, &pokemon.Defm, &pokemon.Atk, &pokemon.Spatk, &pokemon.Speed)

	if err != nil {
		http.Error(w, "Error in pokemon scan", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(pokemon)
	if err != nil {
		http.Error(w, "Error in pokemon encoding", http.StatusInternalServerError)
		return
	}
}

func GetPokemons(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	db := database.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM pokemons")
	if err != nil {
		http.Error(w, "Error in SQL statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var pokemons []models.Pokemon

	rows, err := stmt.Query()
	if err != nil {
		http.Error(w, "Error in SQL Query", http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var pokemon models.Pokemon
		err := rows.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Hp, &pokemon.Def, &pokemon.Defm, &pokemon.Atk, &pokemon.Spatk, &pokemon.Speed)
		if err != nil {
			http.Error(w, "Error in SQL Scan", http.StatusInternalServerError)
			return
		}
		pokemons = append(pokemons, pokemon)
	}

	responseJson, err := json.Marshal(pokemons)
	if err != nil {
		http.Error(w, "Error while encoding json", http.StatusInternalServerError)
		return
	}

	w.Write(responseJson)
}
