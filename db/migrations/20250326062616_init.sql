-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" SCHEMA public;
CREATE TABLE IF NOT EXISTS blog_tag (
    id uuid PRIMARY KEY DEFAULT public.uuid_generate_v4(),
    name VARCHAR(100) DEFAULT '',
    created_by VARCHAR(100) DEFAULT '',
    updated_by VARCHAR(100),
    deleted_by VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE blog_tag;
-- +goose StatementEnd
