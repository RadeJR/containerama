-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id integer primary key autoincrement,
    username text NOT NULL,
    first_name text NOT NULL,
    last_name text,
    email text,
    password_hash text NOT NULL,
    role text NOT NULL,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
