package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/nicolasrsaraiva/database"
	"github.com/nicolasrsaraiva/models"
)

func CreateTrainer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var trainer models.Trainer

	err := json.NewDecoder(r.Body).Decode(&trainer)
	if err != nil {
		http.Error(w, "Invalid json data", http.StatusBadRequest)
		return
	}

	db := database.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO trainers (name, region) VALUES ($1, $2)")
	if err != nil {
		log.Fatalf("Error preparing SQL statement: %s", err)
		return
	}

	_, err = stmt.Exec(trainer.Name, trainer.Region)
	if err != nil {
		log.Fatalf("Error executing SQL statement: %s", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetTrainer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	trainerId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid trainer id", http.StatusBadRequest)
	}

	trainer, err := getTrainerByID(trainerId)

	if err != nil {
		http.Error(w, "Invalid trainer id", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trainer)
}

func GetTrainers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db := database.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM trainers")
	if err != nil {
		http.Error(w, "Error in SQL statement prepare", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		http.Error(w, "Error in SQL query execution", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var trainers []models.Trainer

	for rows.Next() {
		var trainer models.Trainer
		err := rows.Scan(&trainer.Id, &trainer.Name, &trainer.Region, &trainer.PokemonsId)

		if err != nil {
			http.Error(w, "Error while scanning row", http.StatusInternalServerError)
			return
		}

		trainers = append(trainers, trainer)
	}
	responseJson, err := json.Marshal(trainers)
	if err != nil {
		http.Error(w, "Error while encoding json", http.StatusInternalServerError)
	}
	//w.Header().Set("Content-Type", "application/json")
	w.Write(responseJson)
}

func getTrainerByID(id int) (*models.Trainer, error) {
	db := database.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM trainers WHERE ID=$1")
	if err != nil {
		log.Fatalf("Error preparing SQL statement: %s", err)
	}

	var trainer models.Trainer
	row := stmt.QueryRow(id)

	err = row.Scan(&trainer.Id, &trainer.Name, &trainer.Region, &trainer.PokemonsId)

	if err != nil {
		return nil, err
	}

	return &trainer, err
}

func CreateTeamIncomplete(w http.ResponseWriter, r *http.Request) { //incomplete
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db := database.ConnectDB()
	defer db.Close()

	var pokemons []string
	var poketeam []models.Pokemon

	err := json.NewDecoder(r.Body).Decode(&pokemons)
	if err != nil {
		http.Error(w, "Error decoding JSON to array", http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare("SELECT * FROM pokemons WHERE NAME=$1")
	if err != nil {
		http.Error(w, "Error in SQL prepare", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	for _, pokemonName := range pokemons {

		var pkm models.Pokemon

		row := stmt.QueryRow(pokemonName)

		err = row.Scan(&pkm.Id, &pkm.Name, &pkm.Hp, &pkm.Def, &pkm.Defm, &pkm.Atk, &pkm.Spatk, &pkm.Speed)
		if err != nil {
			http.Error(w, "Error in row scan", http.StatusInternalServerError)
			return
		}

		poketeam = append(poketeam, pkm)
	}

	responseJson, err := json.Marshal(poketeam)
	if err != nil {
		http.Error(w, "Error marshaling json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db := database.ConnectDB()
	defer db.Close()

	var idPokemons [6]int
	err := json.NewDecoder(r.Body).Decode(&idPokemons)

	if err != nil {
		http.Error(w, "Error decoding json", http.StatusInternalServerError)
		return
	}

	for _, id := range idPokemons {
		if id <= 0 || id > 20 {
			http.Error(w, "Pokemons doesnt exists", http.StatusBadRequest)
		}
	}

	idTrainer := r.URL.Query().Get("id")

	stmt, err := db.Prepare("UPDATE trainers SET team = ARRAY[$1, $2, $3, $4, $5, $6] where id=$7")
	if err != nil {
		http.Error(w, "Error in prepare sql statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(idPokemons[0], idPokemons[1], idPokemons[2], idPokemons[3], idPokemons[4], idPokemons[5], idTrainer)
	if err != nil {
		http.Error(w, "Error in sql execution", http.StatusInternalServerError)
		return
	}

}
