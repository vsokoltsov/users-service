-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table users(
    id serial primary key,
    email varchar(255) not null,
    first_name varchar,
    last_name varchar,
    created_at timestamp,
    updated_at timestamp
);
create unique index users_email on users (email);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table users;