// TODO: quando houver chamada de busca em branco, tratar o retorno de erro

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

func (repositorio repositorioDb) BuscarId(usuarioId uint64) (models.Usuario, error) {
	row, err := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = $1", usuarioId)
	if err != nil {
		return models.Usuario{}, err
	}
	defer row.Close()

	var usuario models.Usuario

	if row.Next() {
		if err = row.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

func (repositorio repositorioDb) Atualizar(usuarioId uint64, usuario models.Usuario) error {
	statement, err := repositorio.db.Prepare(
		"update usuarios set nome = $1, nick = $2, email = $3 where id = $4",
	)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuarioId); err != nil {
		return err
	}

	return nil
}

func (repositorio repositorioDb) Deletar(usuarioId uint64) error {
	statement, err := repositorio.db.Prepare(
		"delete from usuarios where id = $1",
	)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuarioId); err != nil {
		return err
	}

	return nil
}

func (repositorio repositorioDb) BuscarPorEmail(email string) (models.Usuario, error) {
	row, err := repositorio.db.Query("select id, senha from usuarios where email = $1", email)
	if err != nil {
		return models.Usuario{}, err
	}
	defer row.Close()

	var usuario models.Usuario

	if row.Next() {
		if err = row.Scan(
			&usuario.Id,
			&usuario.Senha,
		); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

// Permite um usuario seguir o outro
func (repositorio repositorioDb) Seguir(usuarioId, seguidorId uint64) error {
	statement, err := repositorio.db.Prepare(
		"insert into seguidores (usuario_id, seguidor_id) values ($1,$2) on conflict do nothing",
	)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuarioId, seguidorId); err != nil {
		return err
	}

	return nil
}
