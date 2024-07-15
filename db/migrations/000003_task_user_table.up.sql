CREATE TABLE IF NOT EXISTS user_task (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL ,
    task_id INT NOT NULL
);

CREATE INDEX idx_user_id ON  user_task(user_id);
CREATE INDEX idx_task_id ON  user_task(task_id);