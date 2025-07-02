-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS blog_tag (
    id INTEGER auto_increment NOT NULL,
    name VARCHAR(100) DEFAULT '',
    created_by VARCHAR(100) DEFAULT '',
    updated_by VARCHAR(100),
    deleted_by VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT blog_tags_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE blog_tag;
-- +goose StatementEnd
