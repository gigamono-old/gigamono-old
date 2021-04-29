-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE x_users_workspaces (
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id uuid REFERENCES users(id),
    workspace_id uuid REFERENCES workspaces(id),
    PRIMARY KEY (user_id, workspace_id)
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE x_users_workspaces;
