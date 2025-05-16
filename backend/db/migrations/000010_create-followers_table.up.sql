CREATE TABLE followers (
    follower_id INTEGER NOT NULL,
    followed_id INTEGER NOT NULL,
    PRIMARY KEY(followed_id, follower_id),
    FOREIGN KEY(followed_id) REFERENCES users(id),
    FOREIGN KEY(follower_id) REFERENCES users(id)
)