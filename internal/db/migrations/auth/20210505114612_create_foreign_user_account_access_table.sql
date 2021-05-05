-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE foreign_user_account_access (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id uuid,
    integration_id uuid,
    enc_access_token text,
    enc_refresh_token text
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE foreign_user_account_access;
