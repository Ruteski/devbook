package repositorios

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type repositorioDb struct {
	db *sql.DB
}

func NovoRepositorioUsuarios(db *sql.DB) *repositorioDb {
	return &repositorioDb{db}
}

// METODO criar usuario
func (repositorio repositorioDb) Criar(usuario models.Usuario) (uint64, error) {
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

func (repositorio repositorioDb) Buscar(nomeNick string) ([]models.Usuario, error) {
	nomeNick = fmt.Sprintf("%%%s%%", nomeNick) // retorna %algumaCoisa%

	rows, err := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios where nome ilike $1 or nick like $2", nomeNick, nomeNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []models.Usuario

	for rows.Next() {
		var usuario models.Usuario

		if err = rows.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}
