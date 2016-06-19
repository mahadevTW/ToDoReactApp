
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE to_do_list (
    id          integer PRIMARY KEY,
    text        text NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE to_do_list;