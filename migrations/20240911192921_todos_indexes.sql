-- +goose Up
-- +goose StatementBegin
create index todos_created_at_index ON todos(created_at);
create index todos_deleted_at_index ON todos(deleted_at);
create index todos_composite_created_at_deleted_at on todos(created_at, deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists todos_created_at_index;
drop index if exists todos_deleted_at_index;
drop index if exists todos_composite_created_at_deleted_at;
-- +goose StatementEnd
