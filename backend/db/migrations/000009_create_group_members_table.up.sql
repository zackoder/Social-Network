CREATE TABLE group_members (
    user_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    role TEXT NOT NULL CHECK (role IN ('creator', 'member')),
    PRIMARY KEY(user_id, group_id),
    FOREIGN KEY(group_id) REFERENCES groups(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
);
