CREATE TABLE events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    group_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    event_time DATETIME NOT NULL,
    created_by INTEGER NOT NULL,
    FOREIGN KEY(group_id) REFERENCES groups(id),
    FOREIGN KEY(created_by) REFERENCES users(id)
);
