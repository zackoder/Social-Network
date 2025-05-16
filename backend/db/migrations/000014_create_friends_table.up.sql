CREATE TABLE friends (
    post_id INTEGER NOT NULL ,
    friend_id INTEGER NOT NULL ,
    PRIMARY KEY (post_id,friend_id),
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (friend_id) REFERENCES users(id)
);