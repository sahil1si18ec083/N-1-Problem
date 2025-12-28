package main

import (
	"database/sql"
	"fmt"
)

func SeedData(db *sql.DB) {
	deleteuserQuery := `Delete Table IF EXISTS users`
	_, err := db.Exec(deleteuserQuery)
	if err != nil {
		fmt.Errorf("Error while deleting user table")
	}
	query := `CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`
	_, err = db.Exec(query)
	if err != nil {
		fmt.Errorf("Error while creating user table")
	}

	deletepostQuery := `Delete Table IF EXISTS posts`
	_, err = db.Exec(deletepostQuery)
	if err != nil {
		fmt.Errorf("Error while deleting post table")
	}
	query_posts := `CREATE Table IF NOT EXISTS posts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT not null,
	body TEXT not null,
	userId INTEGER not null,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	Foreign key (userId) REFERENCES Users (id)
	on delete cascade
	on update cascade
    );`
	_, err = db.Exec(query_posts)
	if err != nil {
		fmt.Errorf("Error while creating user table")
	}
	query_insert_users := `INSERT INTO users (name, email) VALUES
('User1', 'user1@gmail.com'),
('User2', 'user2@gmail.com'),
('User3', 'user3@gmail.com'),
('User4', 'user4@gmail.com'),
('User5', 'user5@gmail.com'),
('User6', 'user6@gmail.com'),
('User7', 'user7@gmail.com'),
('User8', 'user8@gmail.com'),
('User9', 'user9@gmail.com'),
('User10', 'user10@gmail.com'),
('User11', 'user11@gmail.com'),
('User12', 'user12@gmail.com'),
('User13', 'user13@gmail.com'),
('User14', 'user14@gmail.com'),
('User15', 'user15@gmail.com'),
('User16', 'user16@gmail.com'),
('User17', 'user17@gmail.com'),
('User18', 'user18@gmail.com'),
('User19', 'user19@gmail.com'),
('User20', 'user20@gmail.com'),
('User21', 'user21@gmail.com'),
('User22', 'user22@gmail.com'),
('User23', 'user23@gmail.com'),
('User24', 'user24@gmail.com'),
('User25', 'user25@gmail.com'),
('User26', 'user26@gmail.com'),
('User27', 'user27@gmail.com'),
('User28', 'user28@gmail.com'),
('User29', 'user29@gmail.com'),
('User30', 'user30@gmail.com'),
('User31', 'user31@gmail.com'),
('User32', 'user32@gmail.com'),
('User33', 'user33@gmail.com'),
('User34', 'user34@gmail.com'),
('User35', 'user35@gmail.com'),
('User36', 'user36@gmail.com'),
('User37', 'user37@gmail.com'),
('User38', 'user38@gmail.com'),
('User39', 'user39@gmail.com'),
('User40', 'user40@gmail.com'),
('User41', 'user41@gmail.com'),
('User42', 'user42@gmail.com'),
('User43', 'user43@gmail.com'),
('User44', 'user44@gmail.com'),
('User45', 'user45@gmail.com'),
('User46', 'user46@gmail.com'),
('User47', 'user47@gmail.com'),
('User48', 'user48@gmail.com'),
('User49', 'user49@gmail.com'),
('User50', 'user50@gmail.com'),
('User51', 'user51@gmail.com');`
	_, err = db.Exec(query_insert_users)
	if err != nil {
		fmt.Errorf("Error while creating user table")
	}

	query_insert_posts := `INSERT INTO posts (title, body, userId) VALUES
('Post 1A', 'Content for user 1 - post A', 1),
('Post 1B', 'Content for user 1 - post B', 1),

('Post 2A', 'Content for user 2 - post A', 2),
('Post 2B', 'Content for user 2 - post B', 2),

('Post 3A', 'Content for user 3 - post A', 3),
('Post 3B', 'Content for user 3 - post B', 3),

('Post 4A', 'Content for user 4 - post A', 4),
('Post 4B', 'Content for user 4 - post B', 4),

('Post 5A', 'Content for user 5 - post A', 5),
('Post 5B', 'Content for user 5 - post B', 5),

('Post 6A', 'Content for user 6 - post A', 6),
('Post 6B', 'Content for user 6 - post B', 6),

('Post 7A', 'Content for user 7 - post A', 7),
('Post 7B', 'Content for user 7 - post B', 7),

('Post 8A', 'Content for user 8 - post A', 8),
('Post 8B', 'Content for user 8 - post B', 8),

('Post 9A', 'Content for user 9 - post A', 9),
('Post 9B', 'Content for user 9 - post B', 9),

('Post 10A', 'Content for user 10 - post A', 10),
('Post 10B', 'Content for user 10 - post B', 10),

('Post 11A', 'Content for user 11 - post A', 11),
('Post 11B', 'Content for user 11 - post B', 11),

('Post 12A', 'Content for user 12 - post A', 12),
('Post 12B', 'Content for user 12 - post B', 12),

('Post 13A', 'Content for user 13 - post A', 13),
('Post 13B', 'Content for user 13 - post B', 13),

('Post 14A', 'Content for user 14 - post A', 14),
('Post 14B', 'Content for user 14 - post B', 14),

('Post 15A', 'Content for user 15 - post A', 15),
('Post 15B', 'Content for user 15 - post B', 15),

('Post 16A', 'Content for user 16 - post A', 16),
('Post 16B', 'Content for user 16 - post B', 16),

('Post 17A', 'Content for user 17 - post A', 17),
('Post 17B', 'Content for user 17 - post B', 17),

('Post 18A', 'Content for user 18 - post A', 18),
('Post 18B', 'Content for user 18 - post B', 18),

('Post 19A', 'Content for user 19 - post A', 19),
('Post 19B', 'Content for user 19 - post B', 19),

('Post 20A', 'Content for user 20 - post A', 20),
('Post 20B', 'Content for user 20 - post B', 20),

('Post 21A', 'Content for user 21 - post A', 21),
('Post 21B', 'Content for user 21 - post B', 21),

('Post 22A', 'Content for user 22 - post A', 22),
('Post 22B', 'Content for user 22 - post B', 22),

('Post 23A', 'Content for user 23 - post A', 23),
('Post 23B', 'Content for user 23 - post B', 23),

('Post 24A', 'Content for user 24 - post A', 24),
('Post 24B', 'Content for user 24 - post B', 24),

('Post 25A', 'Content for user 25 - post A', 25),
('Post 25B', 'Content for user 25 - post B', 25),

('Post 26A', 'Content for user 26 - post A', 26),
('Post 26B', 'Content for user 26 - post B', 26),

('Post 27A', 'Content for user 27 - post A', 27),
('Post 27B', 'Content for user 27 - post B', 27),

('Post 28A', 'Content for user 28 - post A', 28),
('Post 28B', 'Content for user 28 - post B', 28),

('Post 29A', 'Content for user 29 - post A', 29),
('Post 29B', 'Content for user 29 - post B', 29),

('Post 30A', 'Content for user 30 - post A', 30),
('Post 30B', 'Content for user 30 - post B', 30),

('Post 31A', 'Content for user 31 - post A', 31),
('Post 31B', 'Content for user 31 - post B', 31),

('Post 32A', 'Content for user 32 - post A', 32),
('Post 32B', 'Content for user 32 - post B', 32),

('Post 33A', 'Content for user 33 - post A', 33),
('Post 33B', 'Content for user 33 - post B', 33),

('Post 34A', 'Content for user 34 - post A', 34),
('Post 34B', 'Content for user 34 - post B', 34),

('Post 35A', 'Content for user 35 - post A', 35),
('Post 35B', 'Content for user 35 - post B', 35),

('Post 36A', 'Content for user 36 - post A', 36),
('Post 36B', 'Content for user 36 - post B', 36),

('Post 37A', 'Content for user 37 - post A', 37),
('Post 37B', 'Content for user 37 - post B', 37),

('Post 38A', 'Content for user 38 - post A', 38),
('Post 38B', 'Content for user 38 - post B', 38),

('Post 39A', 'Content for user 39 - post A', 39),
('Post 39B', 'Content for user 39 - post B', 39),

('Post 40A', 'Content for user 40 - post A', 40),
('Post 40B', 'Content for user 40 - post B', 40),

('Post 41A', 'Content for user 41 - post A', 41),
('Post 41B', 'Content for user 41 - post B', 41),

('Post 42A', 'Content for user 42 - post A', 42),
('Post 42B', 'Content for user 42 - post B', 42),

('Post 43A', 'Content for user 43 - post A', 43),
('Post 43B', 'Content for user 43 - post B', 43),

('Post 44A', 'Content for user 44 - post A', 44),
('Post 44B', 'Content for user 44 - post B', 44),

('Post 45A', 'Content for user 45 - post A', 45),
('Post 45B', 'Content for user 45 - post B', 45),

('Post 46A', 'Content for user 46 - post A', 46),
('Post 46B', 'Content for user 46 - post B', 46),

('Post 47A', 'Content for user 47 - post A', 47),
('Post 47B', 'Content for user 47 - post B', 47),

('Post 48A', 'Content for user 48 - post A', 48),
('Post 48B', 'Content for user 48 - post B', 48),

('Post 49A', 'Content for user 49 - post A', 49),
('Post 49B', 'Content for user 49 - post B', 49),

('Post 50A', 'Content for user 50 - post A', 50),
('Post 50B', 'Content for user 50 - post B', 50);
`
	_, err = db.Exec(query_insert_posts)
	if err != nil {
		fmt.Errorf("Error while creating user table")
	}
}
