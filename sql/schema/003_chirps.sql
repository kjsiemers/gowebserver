-- +goose Up
CREATE TABLE users (
id UUID primary key,
created_at TIMESTAMP not null,
updated_at TIMESTAMP not null,
email TEXT not null unique
);
CREATE TABLE chirps_table (
id UUID primary key,
created_at TIMESTAMP not null,
updated_at TIMESTAMP not null,
body TEXT not null,
user_id UUID not null, 
foreign key (user_id) references users (id) on delete cascade
);
-- +goose Down
DROP TABLE users cascade;
DROP TABLE chirps_table;
