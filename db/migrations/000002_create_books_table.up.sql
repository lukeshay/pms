CREATE TABLE books (
  id TEXT PRIMARY KEY,
  user_id TEXT NOT NULL,
  title TEXT NOT NULL,
  author TEXT NOT NULL,
  rating INTEGER,
  purchased_at TIMESTAMP,
  finished_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL,
  created_by TEXT NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by TEXT NOT NULL,
  deleted_at TIMESTAMP,
  deleted_by TEXT,
  FOREIGN KEY(user_id) REFERENCES users(id)
);