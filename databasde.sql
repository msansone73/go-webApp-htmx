/*
crie um script para criação de um banco de dados em postgresql que tenha:
tabela users com os campos:
id primary key e autoincrement
name varchar
email varchar unico
password varchar
enabed boolean
created_at datetime
updated_at datetime

tabela stocks com os campos:
id primary key e autoincrement
code varchar
name varchar
type stock or fii 
created_at datetime
updated_at datetime

tabela transactions com os campos:
id primary key e autoincrement
user userid Foreng key
stock stockid Foreng key
type buy or sell
value decimal
quantity integer
data_at datetime

tabela earnings com os campos:
id primary key e autoincrement
user userid Foreng key
stock stockid Foreng key
type dividends or JCP
value decimal
quantity integer
data_at datetime



by Chat GPT:
*/

-- Criação da tabela 'users'
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) not null,
    email VARCHAR(255) UNIQUE not null,
    password VARCHAR(255) not null,
    enabled BOOLEAN default true,
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now()
);

CREATE TABLE stock_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) not null,
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now()
);
CREATE TABLE stock_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) not null,
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now()
);
-- Criação da tabela 'stocks'
CREATE TABLE stocks (
    code VARCHAR(20)  PRIMARY KEY,
    name VARCHAR(255) not null,
    type int REFERENCES stock_types(id),
    category int REFERENCES stock_categories(id),
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now()
);

-- Criação da tabela 'transactions'
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    stock_code VARCHAR(20) REFERENCES stocks(code),
    type VARCHAR(1) CHECK (type IN ('B', 'S')) not null,
    value DECIMAL(10, 2) not null,
    quantity INT not null,
    date TIMESTAMP not null
);

-- Criação da tabela 'earnings'
CREATE TABLE income (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    stock_code VARCHAR(20) REFERENCES stocks(code),
    value DECIMAL(10, 2) not null,
    quantity INT not null,  
    date TIMESTAMP not null
);

insert into users (name, email, password) values ('Marcio Sansone', 'msansone@gmail.com', 'senha');

insert into stock_types (name) values ('stock');
insert into stock_types (name) values ('fii');

insert into stocks  (code, name, type) values ('BBAS3', 'Banco do Brasil SA', 1);
insert into stocks  (code, name, type) values ('ITSA3', 'Itausa SA', 1);
insert into stocks  (code, name, type) values ('PETR4', 'Petroleo Brasileiro SA Petrobras Preference Shares', 1);
insert into stocks  (code, name, type) values ('SANB3', 'Banco Santander Brasil SA', 1);
insert into stocks  (code, name, type) values ('TAEE4', 'Transmissora Alianca de Enrga Eltrca S/A Preference Shares', 1);
insert into stocks  (code, name, type) values ('CXSE3', 'Caixa Seguridade Participacoes SA', 1);
insert into stocks  (code, name, type) values ('KLBN4', 'Klabin SA Preference Shares', 1);
insert into stocks  (code, name, type) values ('TRPL4', 'CTEEP Cmpnh d Trnsmss d nrg ltrc Plst Preference Shares', 1);
insert into stocks  (code, name, type) values ('GGBR4', 'Gerdau SA Preference Shares', 1);
insert into stocks  (code, name, type) values ('FESA4', 'Companhia de Ferro Ligas da Bah Frbs Preference Shares', 1);
insert into stocks  (code, name, type) values ('GOAU3', 'Metalurgica Gerdau SA', 1);
insert into stocks  (code, name, type) values ('VIVT3', 'Telefonica Brasil SA', 1);
insert into stocks  (code, name, type) values ('CMIG4', 'Energy of Minas Gerais Co Preference Shares', 1);
insert into stocks  (code, name, type) values ('CSMG3', 'Companhia de Saneamento d Mns Grs CPS MG', 1);
insert into stocks  (code, name, type) values ('SAPR11 ', 'Companhia de Saneamento Parana SANEPAR Unit', 1);
insert into stocks  (code, name, type) values ('TASA4', 'Taurus Armas SA Preference Shares', 1);
insert into stocks  (code, name, type) values ('BBSE3', 'BB Seguridade Participacoes SA', 1);
insert into stocks  (code, name, type) values ('HGRU11', 'CSHG Renda Urbana-FI Imobiliario-FII - Un', 2);
insert into stocks  (code, name, type) values ('HGRE11', 'CSHG Real Estate FI Imobiliario-FII - Un', 2);
insert into stocks  (code, name, type) values ('XPLG11', 'XP Log Fundo de Investimento Imobiliario-FII', 2);
insert into stocks  (code, name, type) values ('RBRR11', 'FI Imobiliario RBR Rendimento High Grade', 2);
insert into stocks  (code, name, type) values ('CPTS11', 'Capitania Securities II Fundo de Investimento Imobiliario Closed Fund', 2);
insert into stocks  (code, name, type) values ('MXRF11', 'Maxi Renda Fundo de Investimento Imobiliario - FII', 2);

insert into transactions (user_id, stock_code, type, value, quantity, date) values (1, 'BBAS3', 'B', 30.00, 100, '2021-01-01');

insert into income (user_id, stock_code, value, quantity, date) values (1, 'BBAS3', 0.50, 100, '2021-01-01');

update transactions set value = 32.00 where id = 1;