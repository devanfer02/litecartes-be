CREATE TABLE IF NOT EXISTS completed_task (
    id  INTEGER PRIMARY KEY AUTO_INCREMENT,
    user_uid VARCHAR(255) REFERENCES user(uid),
    task_uid VARCHAR(255) REFERENCES task(uid),
    completed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_uid, task_uid)
) Engine = InnoDB