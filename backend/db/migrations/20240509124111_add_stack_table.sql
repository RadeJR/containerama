-- +goose Up
-- +goose StatementBegin
CREATE TABLE stacks (
    id integer primary key autoincrement,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp,
    name text NOT NULL,
    path_to_file text NOT NULL,
    webhook text,
    branch text,
    user_id text NOT NULL,
    repository_id int,
    FOREIGN KEY(repository_id) REFERENCES repositories(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stacks;
-- +goose StatementEnd
