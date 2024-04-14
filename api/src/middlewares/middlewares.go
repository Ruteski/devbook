package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// escreve informacoes da requisicao no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Host, r.Method, r.RequestURI)
		next(w, r)
	}
}

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando ...")
		next(w, r)
	}
}
