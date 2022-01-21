
create database db_register;

use db_register;


create table permissions (
	id int unsigned auto_increment primary key,
	`name` varchar(100) not null,
    `created_at` datetime default now()
);

INSERT INTO permissions ( name ) VALUES ('USER'), ('ADMIN'), ('DIRECTOR'),('BI'), ('EMPLOYER');

create table menus (
	id int unsigned auto_increment primary key,
    name varchar(100) not null,
    icon_menu varchar(100),
    url_menu varchar(100) not null,
	`created_at` datetime default now(),
	id_permission int unsigned not null,
	FOREIGN KEY (id_permission) REFERENCES permissions(id)
);

insert into menus (name, id_permission, icon_menu, url_menu) values
	('Novo Pedido', 1, 'bxs-truck', 'novoPedido'),
    ('Acompanhar', 1, 'bxs-map', 'acompanhamento'),
	('Historico', 1, 'bxs-timer', 'historico'),
	('Extrato de Notas', 1, 'bxs-file', 'extratoeNotas'),
    ('Perfil', 1, 'bx-user', 'perfil'),
    ('Cadastrar', 2, null, 'cadastrar'),
    ('Perfil Diretoria', 3, 'bx-user', 'perfil');

create table addresses (
	id int unsigned auto_increment primary key,
	street varchar(255) not null,
    state varchar(100) not null,
    `number` varchar(100) not null,
    country varchar(100) not null,
    complement text,
	`created_at` datetime default now()
);

insert into addresses(street, state, `number`, country) values ('rua padre josé alves', 'CE', '390', 'brazil');

create table users (
	id INT unsigned auto_increment primary key,
    phone varchar(255) not null,
    business varchar(255) not null,
    cpf_cnpj varchar(255) not null,
    email varchar(255) not null,
	`password` varchar(255) not null,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    `created_at` datetime default now(),
	id_address int unsigned not null,
	FOREIGN KEY (id_address) REFERENCES addresses(id)
);

insert into users (phone, business, cpf_cnpj, email, `password`, first_name, last_name, id_address) values 
	('8897613741', 'desenvolvedor', '03506838326', 'higordiegoti@gmail.com', '$2a$04$uxMOqH.TJ58FPe06EnqCHuqnLylq7emaDJ764joC.lxSla.Q0WOd6', 'higor', 'pinheiro', 1);

create table user_permissions(
    id int unsigned auto_increment primary key,
    id_user int unsigned not null,
    id_permission int unsigned not null,
    created_at datetime default now(),
    FOREIGN KEY (id_user) REFERENCES users(id),
    FOREIGN KEY (id_permission) REFERENCES permissions(id)
);

insert into user_permissions (id_user, id_permission) values (1, 1);

select p.id as id, p.name as name from permissions as p
				inner join user_permissions as u on p.id = u.id_permission
				where u.id_user = 1