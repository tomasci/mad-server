-- +goose Up
-- +goose StatementBegin
create table users_todos(
    user_id uuid not null constraint users_todos_user_id_fk_users_id references users(id),
    todo_id uuid not null constraint users_todos_todo_id_fk_todos_id references todos(id),
    created_at timestamp with time zone,
    primary key (user_id, todo_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users_todos;
-- +goose StatementEnd
