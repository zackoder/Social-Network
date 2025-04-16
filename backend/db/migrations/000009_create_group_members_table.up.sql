CREATE TABLE group_members (
    user_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    PRIMARY KEY(user_id, group_id)
);