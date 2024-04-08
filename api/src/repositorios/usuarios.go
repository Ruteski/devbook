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
		"insert into usuarios (nome,nick,email,senha) values ($1,$2,$3,$4)",
	)

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	_, err = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	row, err := repositorio.db.Query("select max(id) from usuarios")
	if err != nil {
		return 0, err
	}

	var id uint64
	if row.Next() {
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil
}
