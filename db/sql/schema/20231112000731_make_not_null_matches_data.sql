-- +goose Up
-- +goose StatementBegin
ALTER TABLE matches
  ALTER COLUMN hometeam SET NOT NULL,
  ALTER COLUMN awayteam SET NOT NULL,
  ALTER COLUMN datetime SET NOT NULL,
  ALTER COLUMN stadium SET NOT NULL,
  ALTER COLUMN status SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Revert the changes if needed
-- Note: This assumes that the columns were nullable before the migration
ALTER TABLE matches
  ALTER COLUMN hometeam DROP NOT NULL,
  ALTER COLUMN awayteam DROP NOT NULL,
  ALTER COLUMN datetime DROP NOT NULL,
  ALTER COLUMN stadium DROP NOT NULL,
  ALTER COLUMN status DROP NOT NULL;
-- +goose StatementEnd
