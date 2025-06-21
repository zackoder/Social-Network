CREATE TABLE
    posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        post_privacy TEXT,
        title VARCHAR(255) NOT NULL,
        content TEXT NOT NULL,
        user_id INTEGER NOT NULL,
        imagePath TEXT,
        group_id INTEGER,
        createdAt INTEGER,
        groupe_id INTEGER, 
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE
    );