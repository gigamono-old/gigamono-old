-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE foreign_dev_integration_creds (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    integration_id uuid,
    specification jsonb
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE foreign_dev_integration_creds;
