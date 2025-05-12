CREATE TABLE groups_chat (
    id INTEGER PRIMARY KEY,
    group_id INTEGER NOT NULL,
    sender_id INTEGER NOT NULL,
    content VARCHAR(255) NOT NULL,
    imagePath TEXT,
    FOREIGN KEY(group_id) REFERENCES groups(id),
    FOREIGN KEY(sender_id) REFERENCES users(id)
)