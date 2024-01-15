CREATE TABLE IF NOT EXISTS friend (
    followed_uid VARCHAR(255) NOT NULL REFERENCES user(uid),
    follower_uid VARCHAR(255) NOT NULL REFERENCES user(uid),
    followed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(followedID, followerID)
) Engine = InnoDB