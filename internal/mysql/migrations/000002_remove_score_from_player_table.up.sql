-- Up migration: Removes the score column from players
ALTER TABLE players DROP COLUMN score;
