CREATE TABLE groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) NOT NULL UNIQUE,
    description VARCHAR(200) NOT NULL,
    group_oner INTEGER NOT NULL,
    FOREIGN KEY (group_oner) REFERENCES users(id)
);