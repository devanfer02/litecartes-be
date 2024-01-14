CREATE TABLE IF NOT EXISTS task (
    uid     VARCHAR(255) PRIMARY KEY,
    level   INTEGER DEFAULT 0, 
    sign    VARCHAR(150) DEFAULT "",        
    level_category_id INTEGER REFERENCES level_category(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(level, level_category_id)
) Engine = InnoDB