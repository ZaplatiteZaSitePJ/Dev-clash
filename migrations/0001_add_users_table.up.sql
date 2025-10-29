CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(32) NOT NULL UNIQUE,
    email VARCHAR(128) NOT NULL CHECK (POSITION('@' IN email) > 1),
    hashed_password VARCHAR(72) NOT NULL,
    participant_times INTEGER DEFAULT 0,
    prize_times INTEGER DEFAULT 0,
    moderator_times INTEGER DEFAULT 0,
    status VARCHAR(128),
    description TEXT
);
