
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE user (
  id INTEGER PRIMARY KEY,
  uuid UNSIGNED BIG INT
);

CREATE UNIQUE INDEX uuid ON user(uuid);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user;
