package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositorios"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usuario models.Usuario
	if err = json.Unmarshal(body, &usuario); err != nil {
		log.Fatal(err)
	}

	db, err := db.Conectar()
	if err != nil {
		log.Fatal(err)
	}

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	repositorio.Criar(usuario)

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuário!"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário!"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário!"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário!"))
}
