-- vim: filetype=SQL
CREATE TABLE users (
	id            SERIAL PRIMARY KEY,
	username      TEXT UNIQUE NOT NULL,
	password_hash TEXT NOT NULL,
	created_at    TIMESTAMPTZTZ NOT NULL DEFAULT CURRENT_TIMESTAMPTZ
);

CREATE TABLE sessions (
	id         SERIAL PRIMARY KEY,
	user_id    BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	token      TEXT NOT NULL,
	expires_at TIMESTAMPTZTZ NOT NULL,
	created_at TIMESTAMPTZTZ NOT NULL DEFAULT CURRENT_TIMESTAMPTZ
);