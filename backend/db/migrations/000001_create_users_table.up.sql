CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nickname TEXT UNIQUE DEFAULT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    age VARCHAR(50),
    gender VARCHAR(20),
    email VARCHAR(100) UNIQUE,
    avatar VARCHAR(255) DEFAULT "",
    password VARCHAR(100),
    AboutMe TEXT DEFAULT ""
);