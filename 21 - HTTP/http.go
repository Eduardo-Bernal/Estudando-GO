package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola mundo"))
}

func usuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ph é gay ate no golang fdskkk"))
}

func main() {
	// HTTP E UM PROTOCOLO DE COMUNICAÇAO - BASE DA COMUNICAÇAO WEB

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuario", usuario)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
