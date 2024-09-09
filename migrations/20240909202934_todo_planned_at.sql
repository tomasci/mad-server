-- +goose Up
-- +goose StatementBegin
alter table if exists todos
    add column planned_at timestamp with time zone;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table if exists todos
    drop column planned_at;
-- +goose StatementEnd
