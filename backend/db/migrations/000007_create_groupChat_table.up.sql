CREATE TABLE groups_chat (
    id INTEGER PRIMARY KEY,
    group_id INTEGER NOT NULL,
    sonder_id INTEGER NOT NULL,
    content VARCHAR(255) NOT NULL
)