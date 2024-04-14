package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario models.Usuario
	if err = json.Unmarshal(body, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar(); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuario.Id, err = repositorio.Criar(usuario)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := db.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuarios, err := repositorio.Buscar(nomeNick)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
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
