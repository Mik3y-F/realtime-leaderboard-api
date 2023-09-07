-- Down migration: Adds back the score column to players
ALTER TABLE players
ADD COLUMN score INT DEFAULT 0 NOT NULL;
