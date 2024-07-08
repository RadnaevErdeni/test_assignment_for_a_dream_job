CREATE TABLE IF NOT EXISTS task (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(100),
    start_time TIMESTAMP ,
    duration INTERVAL GENERATED ALWAYS AS (end_time - start_time) STORED,
    done BOOLEAN NOT NULL ,
    took BOOLEAN NOT NULL ,
    end_time TIMESTAMP ,
    date_create TIMESTAMP 
)