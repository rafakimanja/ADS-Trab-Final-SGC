package main

import (
	"fmt"
	"net/http"
)

var (
	porta = "5000"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello Word em Golang rodando no Jenkins, rodando na porta: "+porta)
}

func main() {
	http.HandleFunc("/", homePage)

	fmt.Printf("Servidor ligado na porta %s", porta)
	http.ListenAndServe(":"+porta, nil)
}