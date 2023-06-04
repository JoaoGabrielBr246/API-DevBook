CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(20) not null unique,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB; 

/* 
Ir no terminal, copiar tudo, e colar quando entrar no mysql (mysql -u golang -p)
->Criar tabela caso não exista
->Se existir, ou não, usar a tabela devbook
->Caso a tabela usuarios exista, deletar completamente
->Criar nova tabela usuarios
->unique -> garante que o conteúdo da coluna (ou combinação de colunas) assume um valor diferente para cada linha da tabela
->timestamp é uma função própria do mysql, mostrando a hora que foi inserido os dados no BD
*/