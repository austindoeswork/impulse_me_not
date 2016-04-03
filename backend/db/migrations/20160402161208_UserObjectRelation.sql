
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE user_object_relation (
  id INTEGER PRIMARY KEY,
  userid int, 
  objectuuid TEXT, 
  purchased bool,
  time DATETIME
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user_object_relation;
