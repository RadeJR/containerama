-- +goose Up
-- +goose StatementBegin
ALTER TABLE stacks
ADD roles text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE stacks
DROP COLUMN roles;
-- +goose StatementEnd
