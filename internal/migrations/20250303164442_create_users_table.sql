-- +goose Up
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(55) UNIQUE,
    age INT
);
-- +goose Down
DROP TABLE users;