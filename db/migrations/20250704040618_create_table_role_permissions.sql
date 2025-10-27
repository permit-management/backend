-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_role_permissions (
    id SERIAL PRIMARY KEY,
    role_id BIGINT UNSIGNED NOT NULL,
    permission_id BIGINT UNSIGNED NOT NULL,
    allow_read BOOL DEFAULT false NOT NULL, 
    allow_create BOOL DEFAULT false NOT NULL,
    allow_update BOOL DEFAULT false NOT NULL,
    allow_delete BOOL DEFAULT false NOT NULL,
    UNIQUE (role_id, permission_id),
    FOREIGN KEY (role_id) REFERENCES tbl_role(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES tbl_permissions(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_role_permissions;
-- +goose StatementEnd
