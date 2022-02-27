package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	
)

type Livro struct{
	Id int `json:"id"`
	Titulo string `json:"titulo"`
	Autor string `json:"autor"`
}

var Livros []Livro = []Livro{
	Livro{
		Id: 1,
		Titulo: "O Guarani",
		Autor: "José de Alencar",
	},
	Livro{
		Id: 2,
		Titulo: "Cazuza",
		Autor: "Viriato Correia",
	},
	Livro{
	Id: 3,
	Titulo: "Dom Casmurro",
	Autor: "Machado de Assis",
	},
}

func main() {
	configurarServidor()
}

func configurarServidor(){
	configurarRotas()

	fmt.Println("Servidor esta rodando na porta 1337")
	log.Fatal(http.ListenAndServe(":1337", nil))
}

func listarLivros(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Livros)
	
}

func cadastrarLivro(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	

	encoder := json.NewEncoder(w)
	encoder.Encode(Livros[0])

}

func rotearLivros(w http.ResponseWriter, r *http.Request ){
	if r.Method ==  "GET" {
	listarLivros(w, r)

	} else if r.Method == "POST" {
		cadastrarLivro(w, r)
	}
}

func configurarRotas(){
	http.HandleFunc("/", rotaPrincipal)
//	http.HandleFunc("/livros", listarLivros)
	http.HandleFunc("/livros", rotearLivros )
}

func rotaPrincipal (w http.ResponseWriter, r *http.Request){fmt.Fprintf(w, "Bem vindo")}