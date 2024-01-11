CREATE TABLE IF NOT EXISTS question_category (
    uid             VARCHAR(255) PRIMARY KEY,
    category_name   VARCHAR(255) NOT NULL UNIQUE,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) Engine = InnoDB