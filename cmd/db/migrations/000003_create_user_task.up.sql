CREATE TABLE IF NOT EXISTS user_task (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL ,
    task_id INT NOT NULL
)