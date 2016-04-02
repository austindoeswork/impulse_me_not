
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE user_object_relation (
  id INTEGER_PRIMARY_KEY,
  user_id int, 
  object_id int, 
  purchased bool,
  time DATETIME
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user_object_relation;
