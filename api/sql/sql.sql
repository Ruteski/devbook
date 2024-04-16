-- Script para criar o banco de dados devbook se ele não existir
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'devbook') THEN
        CREATE DATABASE devbook;
    END IF;
END $$;

-- Script para conectar ao banco de dados devbook (opcional)
\c devbook

-- Script para criar a tabela usuarios se ela não existir
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'usuarios') THEN
        CREATE TABLE public.usuarios (
            id serial4 NOT NULL,
            nome varchar(50) NOT NULL,
            nick varchar(50) NOT NULL,
            email varchar(50) NOT NULL,
            senha varchar(100) NOT NULL,
            criadoem timestamp DEFAULT CURRENT_TIMESTAMP NULL,
            CONSTRAINT usuarios_email_key UNIQUE (email),
            CONSTRAINT usuarios_nick_key UNIQUE (nick),
            CONSTRAINT usuarios_pkey PRIMARY KEY (id)
        );
    END IF;
END $$;

-- Script para criar a tabela publicacoes se ela não existir
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'publicacoes') THEN
        CREATE TABLE public.publicacoes (
            id serial4 NOT NULL,
            titulo varchar(50) NOT NULL,
            conteudo varchar(300) NOT NULL,
            autor_id int4 NOT NULL,
            curtidas int4 DEFAULT 0 NULL,
            criadaem timestamp DEFAULT CURRENT_TIMESTAMP NULL,
            CONSTRAINT publicacoes_pkey PRIMARY KEY (id),
            CONSTRAINT publicacoes_autor_id_fkey FOREIGN KEY (autor_id) REFERENCES public.usuarios(id) ON DELETE CASCADE
        );
    END IF;
END $$;

-- Script para criar a tabela seguidores se ela não existir
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'seguidores') THEN
        CREATE TABLE public.seguidores (
            usuario_id int4 NOT NULL,
            seguidor_id int4 NOT NULL,
            CONSTRAINT seguidores_pkey PRIMARY KEY (usuario_id, seguidor_id),
            CONSTRAINT seguidores_seguidor_id_fkey FOREIGN KEY (seguidor_id) REFERENCES public.usuarios(id) ON DELETE CASCADE,
            CONSTRAINT seguidores_usuario_id_fkey FOREIGN KEY (usuario_id) REFERENCES public.usuarios(id) ON DELETE CASCADE
        );
    END IF;
END $$;


-- ======================================================

-- Zerar a tabela seguidores
DO $$ 
BEGIN
    IF EXISTS (SELECT * FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'seguidores') THEN
        TRUNCATE TABLE public.seguidores CASCADE;
        
        -- Reiniciar a sequência para o campo id
        IF EXISTS (SELECT * FROM pg_sequences WHERE schemaname = 'public' AND sequencename = 'seguidores_id_seq') THEN
            ALTER SEQUENCE public.seguidores_id_seq RESTART WITH 1;
        END IF;
    END IF;
END $$;

-- Zerar a tabela publicacoes
DO $$ 
BEGIN
    IF EXISTS (SELECT * FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'publicacoes') THEN
        TRUNCATE TABLE public.publicacoes CASCADE;
        
        -- Reiniciar a sequência para o campo id
        IF EXISTS (SELECT * FROM pg_sequences WHERE schemaname = 'public' AND sequencename = 'publicacoes_id_seq') THEN
            ALTER SEQUENCE public.publicacoes_id_seq RESTART WITH 1;
        END IF;
    END IF;
END $$;

-- Zerar a tabela usuarios
DO $$ 
BEGIN
    IF EXISTS (SELECT * FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'usuarios') THEN
        TRUNCATE TABLE public.usuarios CASCADE;
        
        -- Reiniciar a sequência para o campo id
        IF EXISTS (SELECT * FROM pg_sequences WHERE schemaname = 'public' AND sequencename = 'usuarios_id_seq') THEN
            ALTER SEQUENCE public.usuarios_id_seq RESTART WITH 1;
        END IF;
        
        -- Remover a restrição de chave única antes de truncar
        ALTER TABLE public.usuarios DROP CONSTRAINT IF EXISTS usuarios_email_key;
        ALTER TABLE public.usuarios DROP CONSTRAINT IF EXISTS usuarios_nick_key;
    END IF;
END $$;

-- Remover os índices das tabelas (caso existam)
DO $$ 
BEGIN
    IF EXISTS (SELECT * FROM pg_indexes WHERE tablename = 'seguidores') THEN
        DROP INDEX IF EXISTS seguidores_seguidor_id_fkey;
        DROP INDEX IF EXISTS seguidores_usuario_id_fkey;
    END IF;
    
    IF EXISTS (SELECT * FROM pg_indexes WHERE tablename = 'publicacoes') THEN
        DROP INDEX IF EXISTS publicacoes_autor_id_fkey;
    END IF;
    
    -- A restrição de chave única foi removida acima, então podemos remover o índice agora
    IF EXISTS (SELECT * FROM pg_indexes WHERE tablename = 'usuarios') THEN
        DROP INDEX IF EXISTS usuarios_email_key;
        DROP INDEX IF EXISTS usuarios_nick_key;
    END IF;
END $$;
