package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Contato struct {
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
}

var contatos = []Contato{
	{"José", "81980000000", "jose@email.com"},
	{"Maria", "8198000001", "maria@email.com"},
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /contatos", getContatos)     // use o padrão "METHOD endpoint" para especificar o método da rota
	router.HandleFunc("GET /contatos/{id}", getContato) // use wildcards com chaves para indicar que você quer capturar um parâmetro de rota
	router.HandleFunc("POST /contatos", createContato)

	http.ListenAndServe(":8080", router)
}

func getContatos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contatos)
}

func getContato(w http.ResponseWriter, r *http.Request) {
	idAsString := r.PathValue("id") // capture o parâmetro de rota (como string)
	id, err := strconv.Atoi(idAsString)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	if id >= len(contatos) || id < 0 {
		w.WriteHeader(400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contatos[id])
}

func createContato(w http.ResponseWriter, r *http.Request) {
	var contato Contato
	err := json.NewDecoder(r.Body).Decode(&contato)
	if err != nil {
		w.WriteHeader(422)
		return
	}
	contatos = append(contatos, contato)
	w.WriteHeader(200)
}
