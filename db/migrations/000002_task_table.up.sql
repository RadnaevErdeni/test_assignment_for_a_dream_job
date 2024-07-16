CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(100),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    duration INTERVAL DEFAULT '00:00'::INTERVAL,
    status VARCHAR(20) CHECK (status IN ('not_started', 'in_progress', 'paused', 'completed')),
    last_resume_time TIMESTAMP,
    last_pause_time TIMESTAMP,
    date_create TIMESTAMP,
    total_pause_duration INTERVAL DEFAULT '00:00'::INTERVAL
);