CREATE TABLE IF NOT EXISTS user (
    uid VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255) NOT NULL, 
    email VARCHAR(255) NOT NULL, 
    subscription_id INTEGER DEFAULT 1,
    school_id INTEGER DEFAULT NULL, 
    total_exp INTEGER DEFAULT 0,
    gems INTEGER DEFAULT 0,
    streaks INTEGER DEFAULT 1,
    last_active DATE DEFAULT CURRENT_TIMESTAMP,
    role VARCHAR(100) DEFAULT "__litecartes-app-user",
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(subscription_id) REFERENCES subscription(id),
    FOREIGN KEY(school_id) REFERENCES school(id)
) Engine = InnoDB
