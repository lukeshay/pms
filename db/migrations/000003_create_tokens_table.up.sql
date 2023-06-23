CREATE TABLE tokens (
  id TEXT PRIMARY KEY,
  user_id TEXT NOT NULL,
  token TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by TEXT NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by TEXT NOT NULL,
  deleted_at TIMESTAMP,
  deleted_by TEXT,
  FOREIGN KEY(user_id) REFERENCES users(id)
);