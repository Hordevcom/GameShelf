-- +goose Up
-- +goose StatementBegin
CREATE TYPE game_status AS ENUM ('completed', 'in_progress', 'abandoned', 'not_owned');

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS games (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    genre TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_games (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    game_id INTEGER NOT NULL REFERENCES games(id) ON DELETE CASCADE,
    game_status game_status NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, game_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_games;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS games;
DROP TYPE IF EXISTS game_status;
-- +goose StatementEnd
