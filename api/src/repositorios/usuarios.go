package repositorios

import (
	"api/src/models"
	"database/sql"
)

type respositorioDb struct {
	db *sql.DB
}

func NovoRepositorioUsuarios(db *sql.DB) *respositorioDb {
	return &respositorioDb{db}
}

// METODO criar usuario
func (repositorio respositorioDb) Criar(usuario models.Usuario) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into usuarios (nome,nick,email,senha) values (?,?,?,?)",
	)

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	id, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}
