-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE preferences (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    details jsonb,
    user_id uuid REFERENCES users(id)
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE preferences;
