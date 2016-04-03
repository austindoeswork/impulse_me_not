
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE object (
  id INTEGER PRIMARY KEY,
  uuid VARCHAR(10), 
  addcount int, 
  purchasecount int
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE object;
