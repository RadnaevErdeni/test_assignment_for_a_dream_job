CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    surname VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL ,
    patronymic VARCHAR(100) NOT NULL ,
    address VARCHAR(100),
    passport_serie INT NOT NULL ,
    passport_number INT NOT NULL
);