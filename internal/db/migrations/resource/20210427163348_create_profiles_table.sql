-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE profiles (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    username text,
    first_name text,
    last_name text,
    email text,
    avatar_32_url text,
    user_id uuid REFERENCES users(id)
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE profiles;
