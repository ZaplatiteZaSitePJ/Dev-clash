CREATE TABLE IF NOT EXISTS users_skills (
    user_id  INTEGER NOT NULL,
    skill_id INTEGER NOT NULL,
    PRIMARY KEY (user_id, skill_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (skill_id) REFERENCES skills(id) ON DELETE CASCADE
);
