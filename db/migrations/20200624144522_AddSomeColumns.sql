
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE labels (
    LABEL            varchar(80),
    DATE            date
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE labels;
