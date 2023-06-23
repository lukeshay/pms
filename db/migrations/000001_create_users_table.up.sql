CREATE TABLE users (
  id TEXT PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  email_verified BOOLEAN NOT NULL,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  password TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by TEXT NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by TEXT NOT NULL,
  deleted_at TIMESTAMP,
  deleted_by TEXT
);