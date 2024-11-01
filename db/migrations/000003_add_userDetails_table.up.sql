CREATE TABLE IF NOT EXISTS user_details(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    balance float DEFAULT 0
);
