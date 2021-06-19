-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE spaces (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    avatar_url text,
    creator_id uuid REFERENCES users(id)
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE spaces;
