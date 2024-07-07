CREATE TABLE task (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(100),
    start_time TIMESTAMP ,
    duration VARCHAR(100),
    done BOOLEAN NOT NULL ,
    took BOOLEAN NOT NULL ,
    end_time TIMESTAMP ,
    date_create TIMESTAMP 
)