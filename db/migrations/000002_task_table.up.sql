CREATE TABLE IF NOT EXISTS task (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(100),
    start_time TIMESTAMP ,
	duration_pause INTERVAL ,
    duration INTERVAL ,
    pause_time TIMESTAMP,
    resume_time TIMESTAMP,
    done BOOLEAN NOT NULL ,
    took BOOLEAN NOT NULL ,
    end_time TIMESTAMP ,
    date_create TIMESTAMP,
    count_pause INTEGER DEFAULT 0
)
