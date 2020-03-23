package services

const usersDDL  =  `create table if not exists users (
    id bigserial primary key,
    name varchar(30) not null,
    surname varchar(30) not null,
    login varchar(30) not null unique,
    password text not null,
    address varchar(30) not null,
    email varchar(30) not null,
    phone varchar(30) not null,
	remove boolean not null default false,
    role_id bigserial not null references roles
);`
const rolesDDL  = `create table if not exists roles (
	id bigserial primary key,
	name varchar(30) not null
);`

