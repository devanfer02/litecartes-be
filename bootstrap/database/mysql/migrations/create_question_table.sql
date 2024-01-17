CREATE TABLE IF NOT EXISTS question (
    uid         VARCHAR(255) PRIMARY KEY,
    category_id VARCHAR(255),
    task_uid    VARCHAR(255) REFERENCES task(uid),
    title       VARCHAR(155),
    literacy    TEXT,
    question    VARCHAR(255),
    answer      VARCHAR(200),
    options     TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    FOREIGN KEY(category_id) REFERENCES question_category(uid)
) Engine = InnoDB