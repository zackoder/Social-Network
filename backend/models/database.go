package models

import (
	"database/sql"
	"fmt"
	"os"
)

func CreateTables() *sql.DB {
	db, err := sql.Open("sqlite3", "social_network.db?_foreign_keys=on")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	query := `
	  CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nickname VARCHAR(50),
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		age VARCHAR(50),
		gender VARCHAR(20),
		email VARCHAR(100) UNIQUE,
		avatar VARCHAR(255),
		password VARCHAR(100)
	  );

	  CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,  -- Add UNIQUE constraint here
		token VARCHAR(255) UNIQUE,
		creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	  );

	  CREATE TABLE IF NOT EXISTS posts (
	 	id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		createdAt INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	  );

	  CREATE TABLE IF NOT EXISTS categories (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name varchar(255) NOT NULL UNIQUE
	  );

	  CREATE TABLE IF NOT EXISTS comments (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		comment TEXT NOT NULL,
		date INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
	  );

	  CREATE TABLE IF NOT EXISTS reactions (
   		id INTEGER PRIMARY KEY AUTOINCREMENT,
  		user_id INTEGER NOT NULL,
   		post_id INTEGER,
   		comment_id INTEGER,
    	reaction_type TEXT NOT NULL,
    	date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    	FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    	FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
    	CHECK (
        	(post_id IS NOT NULL AND comment_id IS NULL) OR
        	(comment_id IS NOT NULL AND post_id IS NULL)
    	)
	  );

	  CREATE TABLE IF NOT EXISTS posts_categories (
		post_id INTEGER NOT NULL,
		category_id INTEGER NOT NULL,
		PRIMARY KEY (post_id,category_id),
		FOREIGN KEY (post_id) REFERENCES posts(id),
		FOREIGN KEY (category_id) REFERENCES categories(id)
	  );
	  
	  CREATE TABLE IF NOT EXISTS messages (
	  	id INTEGER PRIMARY KEY AUTOINCREMENT,
	  	sender_id INTEGER NOT NULL,
		reciever_id INTEGER NOT NULL,
		content TEXT,
		is_read BOOLEAN DEFAULT FALSE,
		creation_date INTEGER,
		FOREIGN KEY (sender_id) REFERENCES users(id),
		FOREIGN KEY (reciever_id) REFERENCES users(id)
	  );

	  CREATE TABLE IF NOT EXISTS groups (
	 	id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(100) NOT NULL,
		description VARCHAR(200) NOT NULL
	  );

	  CREATE TABLE IF NOT EXISTS group_members (
	 	user_id INTEGER NOT NULL,
		group_id INTEGER NOT NULL,
		PRIMARY KEY(user_id, group_id)
	  );
	  `
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
