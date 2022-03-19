drop database db_register;
create database db_register;

use db_register;

create table permissions (
	id int unsigned auto_increment primary key,
	`name` varchar(100) not null,
    `created_at` datetime default now()
);

INSERT INTO permissions ( name ) VALUES ('CLIENT'), ('COLLABORATOR');

create table routes (
    id int unsigned auto_increment primary key,
    `path` text not null,
    `method` varchar(10) not null,
    `created_at` datetime default now(),
    id_permission int unsigned not null,
	FOREIGN KEY (id_permission) REFERENCES permissions(id)
);

insert INTO routes (`path`, method, id_permission) values
      ("/api/users", "GET", 1),
      ("/api/client", "GET", 1),
      ("/api/client/:id", "GET", 1),
      ("/api/destination", "GET", 1),
      ("/api/destination/client/:id", "GET", 1),
      ("/api/destination/:id", "GET", 1),
      ("/api/destination/:id", "DELETE", 1),
      ("/api/deposit", "GET", 1),
      ("/api/deposit/:id", "GET", 1),
      ("/api/deposit/:id", "PUT", 1),
      ("/api/deposit/:id", "DELETE", 1),
      ("/api/deposit", "POST", 1),
      ("/api/carrying", "GET", 1),
      ("/api/carrying/:id", "GET", 1),
      ("/api/carrying/:id", "PUT", 1),
      ("/api/carrying/:id", "DELETE", 1),
      ("/api/product", "POST", 1),
      ("/api/product", "GET", 1),
      ("/api/product/:id", "GET", 1),
      ("/api/product/:id", "PUT", 1),
      ("/api/product/:id", "DELETE", 1),
      ("/api/product", "POST", 1),
      ("/api/truck", "POST", 1),
      ("/api/truck/carry/:id", "GET", 1),
      ("/api/truck/:id", "GET", 1),
      ("/api/truck/:id", "PUT", 1),
      ("/api/truck/:id", "DELETE", 1),
      ("/api/truck", "POST", 1),
      ("/api/driver", "POST", 1),
      ("/api/driver", "GET", 1),
      ("/api/driver/:id", "GET", 1),
      ("/api/driver/:id", "PUT", 1),
      ("/api/driver/:id", "DELETE", 1),
      ("/api/driver", "POST", 1),
      ("/api/logistic/calculate", "POST", 1);

create table addresses (
	id int unsigned auto_increment primary key,
	street varchar(255) not null,
    `state` varchar(100) not null,
    `number` varchar(100) not null,
    country varchar(100) not null,
    neighborhood varchar(200) not null,
    complement text,
	`created_at` datetime default now()
);

insert into addresses(street, state, `number`, neighborhood,  country) values ('rua padre josé alves', 'CE', '390', 'salesianos', 'brazil'), ('rua padre josé alves', 'CE', '390', 'salesianos', 'brazil');


create table users (
	id INT unsigned auto_increment primary key,
    phone varchar(255) not null,
    business varchar(255) not null,
    cpf_cnpj varchar(255) not null,
    forgot text,
    email varchar(255) not null,
	`password` varchar(255) not null,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    code_active varchar(255) null,
    `active` boolean default  1,
    `created_at` datetime default now()
);

create table user_addresses (
    id int unsigned auto_increment primary key,
    id_user int unsigned not null,
    id_address int unsigned not null,
    created_at datetime default now(),
    FOREIGN KEY (id_user) REFERENCES users(id),
    FOREIGN KEY (id_address) REFERENCES addresses(id)
);

insert into users (phone, business, cpf_cnpj, email, `password`, first_name, last_name) values 
	('8897613741', 'desenvolvedor', '03506838326', 'higordiegoti@gmail.com', '$2a$04$SjtifNEohNa/JzxYO3m3E.Vtx/rgqbp7KnYJEzf89yhH.JD1AjY5u', 'higor', 'pinheiro'),
    ('8897613741', 'desenvolvedor', '03506838326', 'colaborator@gmail.com', '$2a$04$SjtifNEohNa/JzxYO3m3E.Vtx/rgqbp7KnYJEzf89yhH.JD1AjY5u', 'colaborator', 'colaborator');

insert into user_addresses (id_user, id_address) values (1, 1), (2, 2);

create table user_permissions(
    id int unsigned auto_increment primary key,
    id_user int unsigned not null,
    id_permission int unsigned not null,
    created_at datetime default now(),
    FOREIGN KEY (id_user) REFERENCES users(id),
    FOREIGN KEY (id_permission) REFERENCES permissions(id)
);

insert into user_permissions (id_user, id_permission) values (1, 1), (2, 2);

