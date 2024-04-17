package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := auth.ExtrairUsuarioId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao models.Publicacao
	if erro := json.Unmarshal(body, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadGateway, erro)
		return
	}

	publicacao.AutorId = usuarioId

	fmt.Println(string(body))

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioPublicacoes(db)
	publicacao.Id, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)

}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
