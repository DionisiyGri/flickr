CREATE TABLE IF NOT EXISTS  images (
	id INTEGER PRIMARY KEY,
	title TEXT,
	link TEXT,
	author_id TEXT,
	vote integer
);

PRAGMA user_version = 000001;