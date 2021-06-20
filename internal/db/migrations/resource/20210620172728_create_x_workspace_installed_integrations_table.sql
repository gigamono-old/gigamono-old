-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE x_workspace_installed_integrations (
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    workspace_id uuid REFERENCES workspaces(id),
    integration_id uuid REFERENCES integrations(id),
    PRIMARY KEY (workspace_id, integration_id)
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE x_workspace_installed_integrations;
