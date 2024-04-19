package repositorios

import (
	"api/src/models"
	"database/sql"
)

type repositorioPublicacaoDb struct {
	db *sql.DB
}

func NovoRepositorioPublicacoes(db *sql.DB) *repositorioPublicacaoDb {
	return &repositorioPublicacaoDb{db}
}

func (repositorio repositorioPublicacaoDb) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values ($1,$2,$3)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	_, err = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorId)
	if err != nil {
		return 0, err
	}

	row, err := repositorio.db.Query("select max(id) from publicacoes")
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

func (repositorio repositorioPublicacaoDb) BuscarPorId(publicacaoId uint64) (models.Publicacao, error) {
	row, erro := repositorio.db.Query(`
		select p.*
		     , u.nick
		  from publicacoes p
		 inner join usuarios u
		         on u.id = p.autor_id
		 where p.id = $1		
	`, publicacaoId)
	if erro != nil {
		return models.Publicacao{}, erro
	}
	defer row.Close()

	var publicacao models.Publicacao

	if row.Next() {
		if erro = row.Scan(
			&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

func (repositorio repositorioPublicacaoDb) Buscar(usuarioId uint64) ([]models.Publicacao, error) {
	rows, erro := repositorio.db.Query(`
		select distinct p.*
		     , u.nick
		  from publicacoes p
		 inner join usuarios u
		         on u.id = p.autor_id
		 inner join seguidores s
		         on s.usuario_id = p.autor_id			
		 where p.autor_id = $1
		    or s.seguidor_id = $2 
		 order by 1 desc
	`, usuarioId, usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var publicacoes []models.Publicacao

	for rows.Next() {
		var publicacao models.Publicacao

		if erro = rows.Scan(
			&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio repositorioPublicacaoDb) Atualizar(publicacaoId uint64, publicacao models.Publicacao) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set titulo = $1, conteudo = $2 where id = $3")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio repositorioPublicacaoDb) Deletar(publicacaoId uint64) error {
	statement, erro := repositorio.db.Prepare("delete from publicacoes where id = $1")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}
