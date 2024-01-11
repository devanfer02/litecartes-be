CREATE TABLE IF NOT EXISTS question (
    uid         VARCHAR(255) PRIMARY KEY,
    category_id VARCHAR(255),
    literacy    VARCHAR(300),
    answer      VARCHAR(200),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    FOREIGN KEY(category_id) REFERENCES question_category(uid)
) Engine = InnoDB