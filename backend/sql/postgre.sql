DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name varchar(50) not null,
    userName varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(20) not null,
    creationDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);