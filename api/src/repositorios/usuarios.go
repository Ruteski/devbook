package repositorios

import (
	"api/src/models"
	"database/sql"
)

type usuarios struct {
	dv *sql.DB
}

func NovoRepositorioUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// METODO criar usuario
func (u usuarios) Criar(usuarios models.Usuario) (uint64, error) {
	return 0, nil
}
