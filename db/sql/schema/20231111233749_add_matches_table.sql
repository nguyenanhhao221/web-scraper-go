-- +goose Up
CREATE TABLE matches (
	id TEXT UNIQUE PRIMARY KEY,
	hometeam TEXT,
	awayteam TEXT,
	datetime TEXT,
	stadium  TEXT,
	status   TEXT
);
-- +goose Down
DROP TABLE matches;
