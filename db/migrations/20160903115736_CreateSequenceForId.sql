
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE SEQUENCE todo_sequence
  start 1
  increment 1;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP SEQUENCE todo_sequence;
