CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    surname VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    patronymic VARCHAR(100) NOT NULL,
    address VARCHAR(100),
    passport_serie INT NOT NULL,
    passport_number INT NOT NULL
);

CREATE INDEX idx_surname ON users(surname);
CREATE INDEX idx_name ON users(name);
CREATE INDEX idx_passport_serie ON users(passport_serie);
CREATE INDEX idx_passport_number ON users(passport_number);
CREATE INDEX idx_patronymic ON users(patronymic);
CREATE INDEX idx_address ON users(address);
