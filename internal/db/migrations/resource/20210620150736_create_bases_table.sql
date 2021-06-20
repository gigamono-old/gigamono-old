-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE bases (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at timestamp with time zone DEFAULT (now() at time zone 'utc'),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    creator_id uuid REFERENCES users(id),
    space_id uuid REFERENCES spaces(id)
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE bases;
