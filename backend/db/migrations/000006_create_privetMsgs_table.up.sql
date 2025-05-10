CREATE TABLE messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sender_id INTEGER NOT NULL,
    reciever_id INTEGER NOT NULL,
    content TEXT,
    imagePath TEXT,
    is_read BOOLEAN DEFAULT FALSE,
    creation_date INTEGER,
    FOREIGN KEY (sender_id) REFERENCES users(id),
    FOREIGN KEY (reciever_id) REFERENCES users(id)
);