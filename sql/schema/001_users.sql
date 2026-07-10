-- +goose Up
CREATE TABLE users (
id integer primary key,
created_at TIMESTAMP not null,
updated_at TIMESTAMP not null,
email TEXT not null unique
);

-- +goose Down
DROP TABLE users;
