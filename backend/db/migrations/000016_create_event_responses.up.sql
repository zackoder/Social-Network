CREATE TABLE event_responses (
    user_id INTEGER NOT NULL,
    event_id INTEGER NOT NULL,
    response TEXT CHECK(response IN ('going', 'not going')),
    PRIMARY KEY (user_id, event_id),
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(event_id) REFERENCES events(id)
);
