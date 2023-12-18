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

-- Criação da tabela 'users'
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    enabled BOOLEAN,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Criação da tabela 'stocks'
CREATE TABLE stocks (
    id SERIAL PRIMARY KEY,
    code VARCHAR(255),
    name VARCHAR(255),
    type VARCHAR(50) CHECK (type IN ('stock', 'fii')),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Criação da tabela 'transactions'
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    stock_id INT REFERENCES stocks(id),
    type VARCHAR(50) CHECK (type IN ('buy', 'sell')),
    value DECIMAL(10, 2),
    quantity INT,
    data_at TIMESTAMP
);

-- Criação da tabela 'earnings'
CREATE TABLE earnings (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    stock_id INT REFERENCES stocks(id),
    type VARCHAR(50) CHECK (type IN ('dividends', 'JCP')),
    value DECIMAL(10, 2),
    quantity INT,
    data_at TIMESTAMP
);
