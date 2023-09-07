-- Up migration: Create score_entries table
CREATE TABLE IF NOT EXISTS score_entries (
    entry_id CHAR(36) NOT NULL PRIMARY KEY,
    player_id CHAR(36) NOT NULL,
    score INT NOT NULL,
    entry_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (player_id) REFERENCES players(id)
);
