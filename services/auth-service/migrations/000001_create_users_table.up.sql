CREATE TABLE IF NOT EXISTS users (
  id          BIGSERIAL PRIMARY KEY,
  name        TEXT NOT NULL,
  password    TEXT,
  email       TEXT UNIQUE NOT NULL,
  age         INT,
	created_at  TIMESTAMP WITHOUT TIME ZONE,
	updated_at  TIMESTAMP WITHOUT TIME ZONE,
	deleted_at  TIMESTAMP WITHOUT TIME ZONE
);

CREATE UNIQUE INDEX idx_users_email ON users (email);