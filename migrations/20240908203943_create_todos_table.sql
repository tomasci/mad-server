-- +goose Up
-- +goose StatementBegin
create table todos(
    id uuid not null primary key,
    title text not null,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists todos;
-- +goose StatementEnd
