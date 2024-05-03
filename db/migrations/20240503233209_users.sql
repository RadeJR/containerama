-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id integer primary key autoincrement,
    username text,
    first_name text,
    last_name text,
    email text,
    password_hash text,
    role text,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
