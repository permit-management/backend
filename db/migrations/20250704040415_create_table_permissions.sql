-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_permissions (
    id SERIAL PRIMARY KEY,
    permission_code VARCHAR(50) UNIQUE NOT NULL, -- MDU
    permission_name VARCHAR(100) NOT NULL, -- Master Data User
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_permissions;
-- +goose StatementEnd
