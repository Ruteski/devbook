insert into usuarios (nome, nick, email, senha) values('usuario1', 'usuario_1', 'usuario1@mail.com', '$2a$10$4bOBuR3Hp9/btsGWAnURHOnb9K7ZBXxsmDPXknFYjLymS5PWXtHHW');
insert into usuarios (nome, nick, email, senha) values('usuario2', 'usuario_2', 'usuario2@mail.com', '$2a$10$4bOBuR3Hp9/btsGWAnURHOnb9K7ZBXxsmDPXknFYjLymS5PWXtHHW');
insert into usuarios (nome, nick, email, senha) values('usuario3', 'usuario_3', 'usuario3@mail.com', '$2a$10$4bOBuR3Hp9/btsGWAnURHOnb9K7ZBXxsmDPXknFYjLymS5PWXtHHW');
insert into usuarios (nome, nick, email, senha) values('usuario4', 'usuario_4', 'usuario4@mail.com', '$2a$10$4bOBuR3Hp9/btsGWAnURHOnb9K7ZBXxsmDPXknFYjLymS5PWXtHHW');
insert into usuarios (nome, nick, email, senha) values('usuario5', 'usuario_5', 'usuario5@mail.com', '$2a$10$4bOBuR3Hp9/btsGWAnURHOnb9K7ZBXxsmDPXknFYjLymS5PWXtHHW');
insert into usuarios (nome, nick, email, senha) values('usuario6', 'usuario_6', 'usuario6@mail.com', '$2a$10$4bOBuR3Hp9/btsGWAnURHOnb9K7ZBXxsmDPXknFYjLymS5PWXtHHW');
insert into usuarios (nome, nick, email, senha) values('usuario7', 'usuario_7', 'usuario7@mail.com', '$2a$10$4bOBuR3Hp9/btsGWAnURHOnb9K7ZBXxsmDPXknFYjLymS5PWXtHHW');
insert into usuarios (nome, nick, email, senha) values('usuario8', 'usuario_8', 'usuario8@mail.com', '$2a$10$4bOBuR3Hp9/btsGWAnURHOnb9K7ZBXxsmDPXknFYjLymS5PWXtHHW');
insert into usuarios (nome, nick, email, senha) values('usuario9', 'usuario_9', 'usuario9@mail.com', '$2a$10$4bOBuR3Hp9/btsGWAnURHOnb9K7ZBXxsmDPXknFYjLymS5PWXtHHW');
insert into usuarios (nome, nick, email, senha) values('usuario10', 'usuario_10', 'usuario19@mail.com', '$2a$10$4bOBuR3Hp9/btsGWAnURHOnb9K7ZBXxsmDPXknFYjLymS5PWXtHHW');


insert into seguidores (usuario_id, seguidor_id) values(1,2);
insert into seguidores (usuario_id, seguidor_id) values(1,3);
insert into seguidores (usuario_id, seguidor_id) values(1,4);
insert into seguidores (usuario_id, seguidor_id) values(2,3);
insert into seguidores (usuario_id, seguidor_id) values(2,4);
insert into seguidores (usuario_id, seguidor_id) values(2,1);
insert into seguidores (usuario_id, seguidor_id) values(3,2);
insert into seguidores (usuario_id, seguidor_id) values(3,1);
insert into seguidores (usuario_id, seguidor_id) values(3,4);
insert into seguidores (usuario_id, seguidor_id) values(5,3);
insert into seguidores (usuario_id, seguidor_id) values(5,1);


insert into publicacoes(titulo, conteudo, autor_id)
values
('Publicação do Usuário 1', 'Essa é a publicação do usuário 1! Oba!', 1),
('Publicação do Usuário 2', 'Essa é a publicação do usuário 2! Oba!', 2),
('Publicação do Usuário 3', 'Essa é a publicação do usuário 3! Oba!', 3);
