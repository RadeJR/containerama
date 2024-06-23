-- +goose Up
-- +goose StatementBegin
CREATE TABLE repositories (
    id integer primary key autoincrement,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp,
    name text NOT NULL,
    encrypted_token text,
    url text NOT NULL,
    user_id text NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE repositories;
-- +goose StatementEnd
